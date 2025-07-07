package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/idtoken"
)

type OAuthService struct {
	GoogleConfig *oauth2.Config
}

func NewOAuthService() *OAuthService {
	return &OAuthService{
		GoogleConfig: &oauth2.Config{
			ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
			Endpoint:     google.Endpoint,
			Scopes:       []string{"openid", "email", "profile"},
		},
	}
}

func (oa *OAuthService) ExchangeGoogleCodeForIDToken(code, redirectURI string) (string, error) {
	ctx := context.Background()

	// Important: set the correct redirect URI for this call
	oa.GoogleConfig.RedirectURL = redirectURI

	token, err := oa.GoogleConfig.Exchange(ctx, code)
	if err != nil {
		return "", fmt.Errorf("failed to exchange code: %w", err)
	}

	// Google returns id_token in the "extra" field
	idTokenRaw := token.Extra("id_token")
	idToken, ok := idTokenRaw.(string)
	if !ok || idToken == "" {
		return "", fmt.Errorf("no id_token in Google response")
	}

	return idToken, nil
}

type GoogleUser struct {
	Email   string
	Name    string
	Picture string
}

func (oa *OAuthService) VerifyGoogleIDToken(idToken string) (*GoogleUser, error) {
	ctx := context.Background()

	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	payload, err := idtoken.Validate(ctx, idToken, clientID)
	if err != nil {
		return nil, fmt.Errorf("invalid ID token: %w", err)
	}

	email, _ := payload.Claims["email"].(string)
	name, _ := payload.Claims["name"].(string)
	picture, _ := payload.Claims["picture"].(string)
	res, _ := json.MarshalIndent(payload.Claims, " ", " ")
	log.Println(string(res))

	return &GoogleUser{
		Email:   email,
		Name:    name,
		Picture: picture,
	}, nil
}

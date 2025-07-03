package service

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"
	"path/filepath"
)

type EmailService struct {
	From        string
	Password    string
	Host        string
	Port        string
	TemplateDir string
}

func NewEmailService() *EmailService {
	return &EmailService{
		From:        os.Getenv("EMAIL_FROM"),
		Password:    os.Getenv("EMAIL_PASSWORD"),
		Host:        os.Getenv("EMAIL_HOST"),
		Port:        os.Getenv("EMAIL_PORT"),
		TemplateDir: "./internal/templates/emails/",
	}
}

func (s *EmailService) parseTemplate(templateName string, params map[string]string) (string, error) {
	tplPath := filepath.Join(s.TemplateDir, templateName)
	log.Println(tplPath)
	tmpl, err := template.ParseFiles(tplPath)
	if err != nil {
		return "", err
	}

	var body bytes.Buffer
	err = tmpl.Execute(&body, params)
	if err != nil {
		return "", err
	}
	return body.String(), nil
}

func (s *EmailService) SendEmail(to, subject, templateName string, params map[string]string) error {
	body, err := s.parseTemplate(templateName, params)
	if err != nil {
		return fmt.Errorf("template error: %w", err)
	}

	// Compose MIME email
	msg := []byte(fmt.Sprintf("Subject: %s\r\n"+
		"MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n%s", subject, body))

	addr := fmt.Sprintf("%s:%s", s.Host, s.Port)
	auth := smtp.PlainAuth("", s.From, s.Password, s.Host)

	return smtp.SendMail(addr, auth, s.From, []string{to}, msg)
}

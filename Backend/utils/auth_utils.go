package utils

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
}
func CheckPassword(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return fmt.Errorf("password does not match")
	}
	return err
}

func GenerateOTP() string {
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

func ValidationErrorToText(err error, req any) []string {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]string, len(ve))
		for i, fe := range ve {
			out[i] = FieldErrorToText(fe, req)
		}
		return out
	}
	return []string{err.Error()}
}

func FieldErrorToText(fe validator.FieldError, req any) string {
	field := jsonFieldName(req, fe.Field())

	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", field, fe.Param())
	case "max":
		return fmt.Sprintf("%s cannot be longer than %s characters", field, fe.Param())
	case "email":
		return fmt.Sprintf("%s must be a valid email", field)
	default:
		return fmt.Sprintf("%s is invalid", field)
	}
}

func jsonFieldName(structType any, fieldName string) string {
	t := reflect.TypeOf(structType)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Name == fieldName {
			tag := f.Tag.Get("json")
			if tag != "" && tag != "-" {
				return strings.Split(tag, ",")[0]
			}
			return fieldName
		}
	}
	return fieldName
}

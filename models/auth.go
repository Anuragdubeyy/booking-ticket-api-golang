package models

import (
	"context"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

type AuthCredentials struct {
	// Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email string `json:"email" validate:"required"`
}

type AuthRepository interface {
	GetOne(ctx context.Context, registerData *AuthCredentials ) (*User, error)
	GetUser(ctx context.Context, query interface{}, args ...interface{}) (*User, error)
}

type AuthService interface {
	Login(ctx context.Context, LoginData *AuthCredentials) (string, *User, error)
	Register(ctx context.Context, registerData *AuthCredentials) (*User, error)
}

// check password matching to hash
 func MatchesHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
 }

//  check if an email is valid
func IsValidEmil(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
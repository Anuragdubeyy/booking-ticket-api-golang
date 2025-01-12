package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Anuragdubeyy/ticket-booking-api/models"
	"github.com/Anuragdubeyy/ticket-booking-api/utils"
	"github.com/gohugoio/hugo/tpl/os"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type AuthService struct {
	repository models.AuthRepository
}

func (s *AuthService) Login(ctx context.Context, loginData *models.AuthCredentials) (string, *models.User, error) {
	user, err := s.repository.GetUser(ctx, "email = ? ", loginData.Email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil, fmt.Errorf("Invalid Credentials")
		}
		return "", nil, err
	}

	if !models.MatchesHash(loginData.Password, user.Password){
		return "", nil, fmt.Errorf("Invalid Credentials")
	}

	claims := jwt.MapClaims{
		"id": user.ID,
		"role": user.Role,
		"exp": time.Now().Add(time.Hour * 168).Unix(),

	}
	
	token ,err := utils.GenerateJWT(claims, jwt.SigningMethodHS256, os.Getenv("JWT_SECRET"))
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
	
}

func (s *AuthService) Register(ctx context.Context, registerData *models.AuthCredentials) (string, *models.User, error) {
	if !models.IsValidEmil(registerData.Email){
		return "",nil, fmt.Errorf("please enter a valid Email")
	}

	if _, err := s.repository.GetUser(ctx, "email = ?", registerData.Email); !errors.Is(err, gorm.ErrRecordNotFound){
		return "", nil, fmt.Errorf("the user emial is already registered")
		
	}
	user, err := s.repository.GetUser(ctx, "email = ? ", registerData.Email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil, fmt.Errorf("Invalid Credentials")
		}
		return "", nil, err
	}

	if !models.MatchesHash(registerData.Password, user.Password){
		return "", nil, fmt.Errorf("Invalid Credentials")
	}

	claims := jwt.MapClaims{
		"id": user.ID,
		"role": user.Role,
		"exp": time.Now().Add(time.Hour * 168).Unix(),

	}
	
	token ,err := utils.GenerateJWT(claims, jwt.SigningMethodHS256, os.Getenv("JWT_SECRET"))
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
	
}

func NewAuthService(repository models.AuthRepository) models.AuthService {
	return &AuthService{
		repository: repository,
	}
}

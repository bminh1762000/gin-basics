package service

import (
	"github.com/bminh1762000/jwt-auth-go/models"
	"github.com/bminh1762000/jwt-auth-go/pkg/repository"
)

type Jwt interface {
	GenerateToken(user models.User) (models.Token, error)
	createRefreshToken(token models.Token) (models.Token, error)
	ValidateToken(accessToken string) (models.User, error)
	ValidateRefreshToken(refreshToken string) (models.User, error)
}

type Authorization interface {
	Login(username, password string) (models.User, error)
	Register(username, password string) (int, error)
}

type Service struct {
	Jwt
	Authorization
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Jwt:           NewJwtService(),
		Authorization: NewAuthService(repo.Authorization),
	}
}

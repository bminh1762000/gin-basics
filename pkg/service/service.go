package service

import "github.com/bminh1762000/jwt-auth-go/models"

type Jwt interface {
	GenerateToken(user models.User) (models.Token, error)
	createRefreshToken(token models.Token) (models.Token, error)
	ValidateToken(accessToken string) (models.User, error)
	ValidateRefreshToken(refreshToken string) (models.User, error)
}

type Service struct {
	Jwt
}

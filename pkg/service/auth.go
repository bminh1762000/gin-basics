package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/bminh1762000/jwt-auth-go/models"
	"github.com/bminh1762000/jwt-auth-go/pkg/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Login(username, password string) (models.User, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (s *AuthService) Register(username, password string) (int, error) {
	var user models.User
	user.Username = username
	user.Password = generatePasswordHash(password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

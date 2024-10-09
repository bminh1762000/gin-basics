package service

import (
	"errors"
	"github.com/bminh1762000/jwt-auth-go/models"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	salt            = "woiroe349@4!@#%$^&*()_+"
	signingKey      = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL        = 12 * time.Hour
	refreshTokenTTL = 74 * time.Hour
)

type JwtService struct {
}

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId string `json:"user_id"`
}

type refreshTokenClaims struct {
	jwt.RegisteredClaims
	Token string `json:"token"`
}

func (j *JwtService) GenerateToken(user models.User) (models.Token, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
		},
		user.Username,
	})

	var err error
	jwtToken := models.Token{}
	jwtToken.AccessToken, err = token.SignedString([]byte(signingKey))
	if err != nil {
		return jwtToken, err
	}

	return j.createRefreshToken(jwtToken)
}

func (j *JwtService) ValidateToken(accessToken string) (models.User, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(signingKey), nil
	})

	var user models.User
	if err != nil {
		return user, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok || !token.Valid {
		return user, errors.New("invalid token")
	}

	user.Username = claims.UserId
	return user, nil
}

func (j JwtService) ValidateRefreshToken(refreshToken string) (models.User, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &refreshTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(signingKey), nil
	})

	var user models.User
	if err != nil {
		return user, err
	}

	payload, ok := token.Claims.(*refreshTokenClaims)
	if !ok || !token.Valid {
		return user, errors.New("invalid token")
	}

	parser := jwt.Parser{}
	accessToken, _, err := parser.ParseUnverified(payload.Token, jwt.MapClaims{})
	if err != nil {
		return user, err
	}

	claims, ok := accessToken.Claims.(*tokenClaims)
	if !ok {
		return user, errors.New("invalid token")
	}

	user.Username = claims.UserId
	return user, nil
}

func (j *JwtService) createRefreshToken(token models.Token) (models.Token, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &refreshTokenClaims{
		jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshTokenTTL)),
		},
		token.AccessToken,
	})

	var err error
	token.RefreshToken, err = refreshToken.SignedString([]byte(signingKey))
	if err != nil {
		return token, err
	}

	return token, nil
}

package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/Demoss/books/internal/domain"
	"github.com/Demoss/books/internal/repository"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	salt      = "sadasdasdasdasffgbvxfb312312"
	signedKey = "kdfgakdsjojtj23423jj423!@#"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user domain.User) error {
	user.Password = generateHashPassword(user.Password)

	return s.repo.CreateUser(user)
}

type claims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generateHashPassword(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signedKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid method")
		}
		return []byte(signedKey), nil
	})
	if err != nil {
		return 0, err
	}

	claim, ok := token.Claims.(*claims)
	if !ok {
		return 0, errors.New("invalid type token")
	}

	return claim.UserId, nil
}

func generateHashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

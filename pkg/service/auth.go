package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/berikarg/portfolio-manager/models"
	"github.com/berikarg/portfolio-manager/pkg/repository"
)

const salt = "irugbN[WIOGH394GREO"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
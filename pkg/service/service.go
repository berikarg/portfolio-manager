package service

import (
	"github.com/berikarg/portfolio-manager/models"
	"github.com/berikarg/portfolio-manager/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type Asset interface {
}

type Saving interface {
}

type Service struct {
	Authorization
	Asset
	Saving
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}

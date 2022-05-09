package service

import "github.com/berikarg/portfolio-manager/pkg/repository"

type Authorization interface {
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
		Asset:         NewAssetService(repos.Asset),
		Saving:        NewSavingService(repos.Saving, repos.Asset),
	}
}

package repository

import (
	"github.com/berikarg/portfolio-manager/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
}

type Asset interface {
}

type Saving interface {
}

type Repository struct {
	Authorization
	Asset
	Saving
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}

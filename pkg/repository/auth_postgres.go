package repository

import (
	"github.com/berikarg/portfolio-manager/models"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	return 0, errors.New("method not implemented")
}

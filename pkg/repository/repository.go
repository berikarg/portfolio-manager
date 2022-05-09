package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
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
		Asset:         NewAssetPostgres(db),
		Saving:        NewSavingPostgres(db),
	}
}

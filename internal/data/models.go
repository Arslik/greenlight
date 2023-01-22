package data

import (
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Movies MovieModel
	Users  UserModel
}

func NewModels(dbpool *pgxpool.Pool) Models {
	return Models{
		Movies: MovieModel{DB: dbpool},
		Users:  UserModel{Pool: dbpool},
	}
}

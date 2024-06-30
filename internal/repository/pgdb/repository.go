package pgdb

import (
	"github.com/jmoiron/sqlx"

	"international_trade/internal/entity"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GetUser(username, password string) (entity.User, error)
}

type HashStorage interface {
	AddingHash(input, typeHash, hash string) (string, error)
	DeleteHash(input, typeHash string) error
}

type PG_Repository struct {
	Authorization
	HashStorage
}

func NewPGRepository(db *sqlx.DB) *PG_Repository {
	return &PG_Repository{
		Authorization: NewAuthPostgres(db),
		HashStorage:   NewHashToPostgres(db),
	}
}

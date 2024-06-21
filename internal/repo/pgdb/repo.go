package pgdb

import (
	"github.com/jmoiron/sqlx"

	"international_trade/internal/entity"
)

const (
	usersTable = "users"
	hashTable = "hash_storage"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GetUser(username, password string) (entity.User, error)
}

type HashStorage interface {
	AddingHash(input, hash string) (string, error)
	// GetAll() ([]autocatalog.Car, error)
	// GetById(carId int) (autocatalog.Car, error)
	// Delete(carId int) error
	// UpdateCatalog(input autocatalog.Car) error
}

type Repository struct {
	Authorization
	HashStorage
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		HashStorage:   NewHashToPostgres(db),
	}
}

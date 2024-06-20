package pgdb

import (
	"github.com/jmoiron/sqlx"

	"international_trade/internal/entity"
)

const (
	usersTable = "users"
	listString = "list_to-hash"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GetUser(username, password string) (entity.User, error)
}

type ListStrings interface {
	// AddingCars(list autocatalog.Car) error
	// GetAll() ([]autocatalog.Car, error)
	// GetById(carId int) (autocatalog.Car, error)
	// Delete(carId int) error
	// UpdateCatalog(input autocatalog.Car) error
}

type Repository struct {
	Authorization
	ListStrings
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		// ListStrings:   NewListCarPostgres(db),
	}
}

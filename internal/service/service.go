package service

import (
    "international_trade/internal/entity"
    repo "international_trade/internal/repo/pgdb"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type ListCars interface {
	// AddingCars(autocatalog.Car) error
	// GetAllCars() ([]autocatalog.Car, error)
	// GetCarById(carId int) (autocatalog.Car, error)
	// Delete(carId int) error
	// UpdateCatalog(input autocatalog.Car) error
}

type Service struct {
	Authorization
	ListCars
}

func NewService(repos *repo.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		// ListCars:      NewListCarsService(repos.ListCars),
	}
}
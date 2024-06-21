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

type ServingString interface {
	CreateNewHash(input string, typeHash string) (string,error)
	GetHash(input string) (string, error)
	DeleteteHash(input string, typeHash string) (string, error)
}

type Service struct {
	Authorization
	ServingString
}

func NewService(repos *repo.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		ServingString: NewHashService(repos.HashStorage),
	}
}

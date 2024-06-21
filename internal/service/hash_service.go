package service

import (
	"international_trade/internal/repo/pgdb"

	. "international_trade/internal/service/processing"
)

type HashService struct {
	repo pgdb.HashStorage
}

func NewHashService(repos pgdb.HashStorage) *HashService {
	return &HashService{repo: repos}
}

func (h *HashService) CreateNewHash(input string, typeHash string) (string, error) {

	hash, err := Processing(input, typeHash)

	if err != nil {
		return "", err
	}

	return h.repo.AddingHash(input, string(hash))
}

package service

import (
	"fmt"
	"international_trade/internal/repo/pgdb"

	"international_trade/internal/repo/redisdb"
	. "international_trade/internal/service/processing"
)

type HashService struct {
	repo pgdb.HashStorage
}

func NewHashService(repos pgdb.HashStorage) *HashService {
	return &HashService{repo: repos}
}

func (h *HashService) CreateNewHash(input string, typeHash string) (string, error) {

	hash, _, err := Processing(input, typeHash)
	if err != nil {
		return "", err
	}

	redisdb.CreateNewEntry(input, typeHash, hash)

	h.repo.AddingHash(input, typeHash, hash)

	return hash, err
}

func (h *HashService) UpdateHash(input string, typeHash string) (string, error) {
	return "", nil
}

func (h *HashService) GetHash(input string) (string, error) {
	return "", nil
}

func (h *HashService) DeleteHash(input string, typeHash string) error {

	err := redisdb.DeleteHash(input, typeHash)
	if err!= nil {
        fmt.Errorf("error deleting hash on Redis: %w", err)
    }
	err = h.repo.DeleteHash(input, typeHash)
	if err!= nil {
        fmt.Errorf("error deleting hash on PG: %w", err)
    }
	return err
}

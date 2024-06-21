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

func (h *HashService) AddingHash(inputKey string, typeHash string) (string, error) {

	hash, err := redisdb.CheckHash(inputKey, typeHash)
	if err != nil {

		hash, _, err = Processing(inputKey, typeHash)
		if err != nil {
			return "", err
		}

		redisdb.CreateNewEntry(inputKey, typeHash, hash)
		h.repo.AddingHash(inputKey, typeHash, hash)
	}

	return hash, err
}
func (h *HashService) GetHash(inputKey, typeHash string) (string, error) {

	hash, err := redisdb.CheckHash(inputKey, typeHash)
	return hash, err
}

func (h *HashService) DeleteHash(inputKey string, typeHash string) error {

	err := redisdb.DeleteHash(inputKey, typeHash)
	if err != nil {
		fmt.Errorf("error deleting hash on Redis: %w", err)
	}
	err = h.repo.DeleteHash(inputKey, typeHash)
	if err != nil {
		fmt.Errorf("error deleting hash on PG: %w", err)
	}
	return err
}

package service

import (
	"fmt"
	"international_trade/internal/repo/pgdb"
	"log"

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

func (h *HashService) UpdateHash(input string, typeHash string) (string, error) {
	return "", nil
}

func (h *HashService) GetHash(input string) (string, error) {
	return "", nil
}

func (h *HashService) DeleteHash(input string, typeHash string) error {

	err := redisdb.DeleteHash(input, typeHash)
	if err != nil {
		fmt.Errorf("error deleting hash on Redis: %w", err)
	}
	err = h.repo.DeleteHash(input, typeHash)
	if err != nil {
		fmt.Errorf("error deleting hash on PG: %w", err)
	}
	return err
}

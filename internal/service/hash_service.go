package service

import (
	"fmt"

	"international_trade/internal/repository/pgdb"
	"international_trade/internal/repository/redisdb"
	. "international_trade/internal/service/processing"
)

type HashService struct {
	repoPG      pgdb.HashStorage
	repoRedis redisdb.RecordStorage
}

func NewHashService(repos pgdb.HashStorage, repoRedis redisdb.RecordStorage) *HashService {
	return &HashService{
		repoPG:      repos,
		repoRedis: repoRedis}
}

// Create a new hash if no existing value is found for the given string. Saving the new value to the database.
func (h *HashService) AddingHash(inputKey string, typeHash string) (string, error) {

	err := checkLengthString(inputKey)
	if err != nil {
		return "", err
	}
	// If the value is not in the database, then create a new one and save it.
	hash, err := h.repoRedis.CheckEntry(inputKey, typeHash)
	if err != nil {

		hash, err = Processing(inputKey, typeHash)
		if err != nil {
			return "", err
		}

		err = h.repoRedis.CreateNewEntry(inputKey, typeHash, hash)
		if err != nil {
			return "", fmt.Errorf("failed to create new entry in Redis: %w", err)
		}
		_, err = h.repoPG.AddingHash(inputKey, typeHash, hash)
		if err != nil {
			return "", fmt.Errorf("failed to create new entry in DB: %w", err)
		}
	}

	return hash, err
}
func (h *HashService) GetHash(inputKey, typeHash string) (string, error) {

	err := checkLengthString(inputKey)
	if err != nil {
		return "", err
	}
	hash, err := h.repoRedis.CheckEntry(inputKey, typeHash)
	return hash, err
}

func (h *HashService) DeleteHash(inputKey string, typeHash string) error {

	err := checkLengthString(inputKey)
	if err != nil {
		return err
	}
	err = h.repoRedis.DeleteEntry(inputKey, typeHash)
	if err != nil {
		fmt.Errorf("error deleting hash on Redis: %w", err)
	}
	err = h.repoPG.DeleteHash(inputKey, typeHash)
	if err != nil {
		fmt.Errorf("error deleting hash on PG: %w", err)
	}
	return err
}

func checkLengthString(inputKey string) error {

	if len(inputKey) > 255 {
		return fmt.Errorf("key is too long")
	}
	return nil
}

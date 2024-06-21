package service

import (
	"international_trade/internal/repo/pgdb"

	. "international_trade/internal/service/processing"
	"international_trade/pkg/redis"
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

	err = redis.RedisClient.Set(input, hash, 0).Err()
	if err != nil {
		return "", err
	}

	_, err = redis.RedisClient.Get(input).Result()
	if err != nil {
		return "", err
	}

	h.repo.AddingHash(input, typeHash, hash)

	return hash, err
}

func (h *HashService) UpdateHash(input string, typeHash string) (string, error) {
	return "", nil
}

func (h *HashService) GetHash(input string) (string, error) {
	return "", nil
}

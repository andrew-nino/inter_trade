package service

import (
	"fmt"
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

	hash, err := Processing(input, typeHash)
	if err != nil {
		return "", err
	}

	err = redis.RedisClient.Set(input, hash, 0).Err()
	if err != nil {
		return "", err
	}

	val, err := redis.RedisClient.Get(input).Result()
	if err != nil {
		return "", err
	}


	h.repo.AddingHash(input, string(hash))

	return fmt.Sprintf("key = %s, value = %s", input, val), err
}

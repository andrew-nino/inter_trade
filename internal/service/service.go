package service

import (
	"international_trade/internal/entity"
	postgres "international_trade/internal/repository/pgdb"
	redis "international_trade/internal/repository/redisdb"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type ServingString interface {
	AddingHash(input, typeHash string) (string, error)
	GetHash(input, typeHash string) (string, error)
	DeleteHash(input, typeHash string) error
}

type Service struct {
	Authorization
	ServingString
}

func NewService(reposPG *postgres.PG_Repository, reposRedis *redis.RedisRepository) *Service {
	return &Service{
		Authorization: NewAuthService(reposPG.Authorization),
		ServingString: NewHashService(reposPG.HashStorage, reposRedis.RecordStorage),
	}
}

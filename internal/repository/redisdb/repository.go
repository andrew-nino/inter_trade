package redisdb

import "github.com/go-redis/redis"

type RecordStorage interface {
	CreateNewEntry(inputKey string, typeHash string, hash string) error
	DeleteEntry(key string, typeHash string) error
	CheckEntry(inputKey string, typeHash string) (string, error)
}

type RedisRepository struct {
	RecordStorage
}

func NewRedisRepository(r *redis.Client) *RedisRepository {
	return &RedisRepository{
		RecordStorage: NewRedisStorage(r),
	}
}

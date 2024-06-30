package redisdb

import (
	"log/slog"
	"os"

	"github.com/go-redis/redis"
)

type RedisStorage struct {
	logger      *slog.Logger
	redisClient *redis.Client
}

func NewRedisStorage(rc *redis.Client) *RedisStorage {
	return &RedisStorage{
		logger:      slog.New(slog.NewJSONHandler(os.Stdout, nil)),
		redisClient: rc,
	}
}

// Creating a new record in the database.
// The full key consists of the input key and the hash type for it.
func (r *RedisStorage) CreateNewEntry(inputKey string, typeHash string, hash string) error {

	fullKey := inputKey + "/" + typeHash

	err := r.redisClient.Set(fullKey, hash, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

// Deleting an existing record from the database.
// The full key consists of the input key and the hash type for it.
func (r *RedisStorage) DeleteEntry(key string, typeHash string) error {

	fullKey := key + "/" + typeHash
	err := r.redisClient.Del(fullKey).Err()
	return err
}

// Checks for the presence of an already stored value.
// The full key consists of the input key and the hash type for it.
func (r *RedisStorage) CheckEntry(inputKey string, typeHash string) (string, error) {

	fullKey := inputKey + "/" + typeHash
	hash, err := r.redisClient.Get(fullKey).Result()

	return hash, err
}

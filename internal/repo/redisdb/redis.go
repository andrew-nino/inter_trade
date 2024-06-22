package redisdb

import (
	. "international_trade/pkg/redis"
)
// Creating a new record in the database.
func CreateNewEntry(inputKey string, typeHash string, hash string) error {

	fullKey := inputKey + "/" + typeHash

	err := RedisClient.Set(fullKey, hash, 0).Err()
	if err != nil {
		return err
	}

	_, err = RedisClient.Get(fullKey).Result()
	if err != nil {
		return err
	}
	return err
}

func DeleteHash(key string, typeHash string) error {

	fullKey := key + "/" + typeHash
	err := RedisClient.Del(fullKey).Err()
	return err
}
// Checks for the presence of an already stored value.
func CheckHash(inputKey string, typeHash string) (string, error) {

	fullKey := inputKey + "/" + typeHash
	hash, err := RedisClient.Get(fullKey).Result()

	return hash, err
}

package redisdb

import (
	. "international_trade/pkg/redis"
)

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

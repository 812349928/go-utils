package redis

import (
	"fmt"
	"github.com/go-redis/redis"
)

func NewQuery(host, port, password, dbname string) (*redis.Client, error) {
	redisClient := &redis.Client{}
	if redisClient != nil {
		return redisClient, nil
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       0,
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		return nil, err
	}

	return redisClient, nil
}

func Set(redisClient *redis.Client, key string, val interface{}, expire int) error {
	if expire > 0 {
		err := redisClient.Do("SET", key, val, "PX", expire).Err()
		if err != nil {
			return err
		}
	} else {
		err := redisClient.Do("SET", key, val).Err()
		if err != nil {
			return err
		}
	}

	return nil
}

func RedisGet(redisClient *redis.Client, key string) (string, error) {
	value, err := redisClient.Do("GET", key).String()
	if err != nil {
		return "", nil
	}

	return value, nil
}

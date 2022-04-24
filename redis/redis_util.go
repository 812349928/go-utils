package redis

import (
	"fmt"
	"github.com/go-redis/redis"
)

type Redis struct {
	Client *redis.Client
}

func NewRedis(host, port, password, dbname string) *Redis {
	redisUtil := &Redis{}

	redisUtil.Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       0,
	})
	_, err := redisUtil.Client.Ping().Result()
	if err != nil {
		return redisUtil
	}

	return redisUtil
}

func (r *Redis) Set(key string, val interface{}, expire int) error {
	if expire > 0 {
		err := r.Client.Do("SET", key, val, "PX", expire).Err()
		if err != nil {
			return err
		}
	} else {
		err := r.Client.Do("SET", key, val).Err()
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Redis) RedisGet(key string) (string, error) {
	value, err := r.Client.Do("GET", key).String()
	if err != nil {
		return "", nil
	}

	return value, nil
}

package redis

import (
	"fmt"
	"github.com/go-redis/redis"
)

type Redis struct {
	Client *redis.Client
}

func (r Redis) Init(host, port, password string) {
	r.Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       0,
	})
	r.Client.Ping().Result()
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

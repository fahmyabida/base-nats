package main

import (
	"sync"
	"time"

	"github.com/go-redis/redis"
)

type redisConfiguration Service

// RedisConfig is var for redisConfiguration
var (
	RedisEngine IRedis
	RedisConfig redisConfiguration
)

// Redis is used to create object redis
type Redis struct {
	Redis *redis.Client
}

func (config redisConfiguration) Initialize() error {
	client := redis.NewClient(&redis.Options{
		Addr:     config.URL,
		Password: "", // no password set
		DB:       0,  // use default DB)
	})
	RedisEngine = Redis{
		Redis: client,
	}
	result := client.Ping()
	return result.Err()
}

func (config redisConfiguration) Job(worker ...func(waitGroup *sync.WaitGroup)) {
	var waitGroup sync.WaitGroup
	for _, work := range worker {
		waitGroup.Add(1)
		go work(&waitGroup)
	}

	waitGroup.Wait()
}

// HMSet is used to create command redis HMSET
func (r Redis) HMSet(key string, field map[string]interface{}) error {
	error := r.Redis.HMSet(key, field)
	return error.Err()
}

// HMGet is used to create command redis HGSET
func (r Redis) HMGet(key string, field ...string) ([]interface{}, error) {
	result, error := r.Redis.HMGet(key, field...).Result()
	return result, error
}

// Set is used to create command redis SET
func (r Redis) Set(key string, field interface{}, timeout time.Duration) error {
	error := r.Redis.Set(key, field, timeout)
	return error.Err()
}

// Get is used to create command redis Get
func (r Redis) Get(key string) (interface{}, error) {
	result, error := r.Redis.Get(key).Result()
	return result, error
}

// Expire is used to create command redis Expire
func (r Redis) Expire(key string, timeout time.Duration) error {
	error := r.Redis.Expire(key, timeout)
	return error.Err()
}

// ExpireAt is used to create command redis Expire
func (r Redis) ExpireAt(key string, timeout time.Time) error {
	error := r.Redis.ExpireAt(key, timeout)
	return error.Err()
}

package main

import (
	"sync"
	"time"

	"github.com/jinzhu/gorm"
)

// IPostgres interface is used to create object postgre
type IPostgres interface {
	DB() *gorm.DB
}

// Engine interface is a blue print to create collection method to provide object to do some task
type Engine interface {
	Initialize() error                             // Stop Function is used  to create some command after the job done
	Job(worker ...func(waitGroup *sync.WaitGroup)) // Job Function is used to run all taks that attach on parameter
}

// IRedis interface is used to create object redis command
type IRedis interface {
	HMSet(key string, field map[string]interface{}) error
	HMGet(key string, field ...string) ([]interface{}, error)
	Set(key string, field interface{}, timeout time.Duration) error
	Get(key string) (interface{}, error)
	Expire(key string, timeout time.Duration) error
	ExpireAt(key string, timeout time.Time) error
}

// INats interface is used to create object nats command
type INats interface {
	Publish(topic string, data interface{})
	Subscribe(topic string, field interface{})
	Request(topic string, field interface{}, reply interface{}, timeout time.Duration)
	PublishRequest(topic string, field interface{}, reply interface{})
}

package main

import (
	"errors"
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//LoadAndInitConfigurationFrom is used to load configuration of natsIO and RedisIO and setup all natsEngine and redisEngine
func LoadAndInitConfigurationFrom(engine Engine) {
	errorStart := engine.Initialize()
	if errorStart != nil {
		Log.Error(errorStart.Error())
	}
	engine.Job(GetConfigurationService)
}

type postgreBKUMaster struct {
	postgre *gorm.DB
}

func (p postgreBKUMaster) DB() *gorm.DB {
	return p.postgre
}

type bkuMasterConfiguration Configuration

// PostgreBKUMaster is instance variable to represent postgre object
var (
	PostgreBKUMaster       IPostgres
	BkuMasterConfiguration bkuMasterConfiguration
)

// New is initilize var BkuMasterConfiguration
func (config bkuMasterConfiguration) New(prop Properties) {
	BkuMasterConfiguration = bkuMasterConfiguration{
		IP:       properties.IP,
		Port:     properties.Port,
		User:     properties.User,
		Password: properties.Password,
		DBName:   properties.DBName,
	}
}

func (config bkuMasterConfiguration) Initialize() error {
	db, err := gorm.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname%s sslmode=disable TimeZone=Asia/Jakarta",
			config.IP,
			config.Port,
			config.User,
			config.Password,
			config.DBName))
	if err != nil {
		return errors.New("Failed to connect database postgres " + config.DBName + " : " + err.Error())
	}
	PostgreBKUMaster = postgreBKUMaster{
		postgre: db,
	}
	return nil
}

func (config bkuMasterConfiguration) Job(worker ...func(waitGroup *sync.WaitGroup)) {
	var wg sync.WaitGroup
	for _, work := range worker {
		wg.Add(1)
		go work(&wg)
	}
	wg.Wait()
}

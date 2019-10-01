package main

import (
	"encoding/json"
	"sync"
	"time"
)

// GetConfigurationService is to initialize Redis and Nats config
func GetConfigurationService(wg *sync.WaitGroup) {
	timeStart := time.Now()
	Log.Event("Starting Load Config" + timeStart.Format("2006-01-02 15:04:05"))
	var (
		modelDBService ServiceConfig
		mapConfig      map[string]interface{}
	)
	PostgreBKUMaster.DB().Exec("SET search_path TO bku_config_ref").Select("properties").
		Where("name = ?", properties.Config).Find(&modelDBService)
	errorParse := json.Unmarshal([]byte(modelDBService.Properties), &mapConfig)
	if errorParse != nil {
		Log.Error("error parsing JSON Redis configuration: " + errorParse.Error())
	}
	// init RedisConfig in redisConfiguration
	RedisConfig = redisConfiguration{
		URL: mapConfig["nats"].(string),
	}
	// init NatsConfig in natsConfiguration
	NatsConfig = natsConfiguration{
		URL: mapConfig["redis"].(string),
	}
	timeFinish := time.Now()
	Log.Event("Finish Load Config" + timeFinish.Format("2006-01-02 15:04:05"))
	wg.Done()
}

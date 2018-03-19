package config

import (
	"encoding/json"
	"os"
	"path"
	"runtime"
)

type Configuration struct {
	MongoURL string
	Database string
}

var globalConfig *Configuration

func GetMongoURL() string {
	if globalConfig == nil {
		globalConfig = loadConfiguration()
	}

	return globalConfig.MongoURL
}

func GetDatabaseName() string {
	if globalConfig == nil {
		globalConfig = loadConfiguration()
	}

	return globalConfig.Database
}

func loadConfiguration() *Configuration {
	env := os.Getenv("APP_ENV")
	var filepath string

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}

	if env == "DEV" || env == "" {
		filepath = path.Dir(filename) + "/dev.json"
	}

	file, _ := os.Open(filepath)
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := Configuration{}
	err := decoder.Decode(&config)

	if err != nil {
		panic(err)
	}

	return &config
}

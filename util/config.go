package util

import (
	"encoding/json"
	"os"
)

var DefaultConfig = map[string]string{
	"host":        "localhost",
	"port":        "8080",
	"remote_host": "localhost",
	"donate_host": "localhost",
}

func InitConfig() {
	config := GetConfig()
	for key, value := range DefaultConfig {
		if _, ok := config[key]; !ok {
			config[key] = value
		}
	}
	SetConfig(config)
}

func CreateConfig() {
	if _, err := os.Stat("config.json"); os.IsNotExist(err) {
		os.Create("config.json")
	}
}

func GetConfig() map[string]string {
	var config map[string]string
	CreateConfig()
	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}
	return config
}

func SetConfig(config map[string]string) {
	CreateConfig()
	file, err := os.OpenFile("config.json", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(config)
	if err != nil {
		panic(err)
	}
}

func SetConfigFromKeyValue(key, value string) {
	config := GetConfig()
	config[key] = value
	SetConfig(config)
}

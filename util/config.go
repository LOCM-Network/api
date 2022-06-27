package util

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var DefaultConfig = map[string]string{
	"host":        "localhost",
	"port":        "8080",
	"remote_host": "localhost",
	"donate_host": "localhost",
}

type Config struct {
	Host       string
	Port       int
	RemoteHost string
	DonateHost string
}

func InitConfig() {
	CreateConfig()
}

func CreateConfig() {
	if _, err := os.Stat("config.json"); os.IsNotExist(err) {
		data := Config{
			Host:       "localhost",
			Port:       8080,
			RemoteHost: "localhost",
			DonateHost: "localhost",
		}
		file, _ := json.MarshalIndent(data, " ", " ")
		_ = ioutil.WriteFile("config.json", file, 0644)
	}
}

func GetConfig() *Config {
	file, _ := ioutil.ReadFile("config.json")
	data := &Config{}
	_ = json.Unmarshal([]byte(file), &data)

	return data
}

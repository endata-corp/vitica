package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"vitica/log"
)

var configFilePath = os.Getenv("VITICA_DIR") + "/config/dev.json"
var config Config

type Config struct {
	DB DBConfig `json:"db"`
}

type DBConfig struct {
	Hostname string `json:"hostname"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	User     string `json:"user"`
	Password string `json:"password"`
	Options  string `json:"options"`
}

func GetDBConfig() string {
	return fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?%v",
		config.DB.User,
		config.DB.Password,
		config.DB.Hostname,
		config.DB.Port,
		config.DB.Database,
		config.DB.Options,
	)
}

func ReadConfig() error {
	jsonConfig, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Error("Error reading config file: %v", configFilePath)
		return err
	}
	json.Unmarshal(jsonConfig, &config)
	return nil
}

func SaveConfig() error {
	jsonConfig, err := json.MarshalIndent(&config, "", "  ")
	if err != nil {
		log.Error("Couldn't marshalled config: %v", err.Error())
		return err
	}
	err = ioutil.WriteFile(configFilePath, jsonConfig, 0644)
	if err != nil {
		log.Error("Error writing config file: %v", err.Error())
		return err
	}
	return nil
}

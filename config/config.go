/*Package config provide global config */
package config

import (
	"encoding/json"
	"os"
)

// GlobalConfig save config from config.json
var GlobalConfig Config

// Config is configuration of this service
type Config struct {
	Host       string
	DbPort     uint `json:"db_port"`
	User       string
	Password   string
	Database   string
	ListenPort uint `json:"listen_port"`
}

// ReadConfig read config.json into global value GlobalConfig
func ReadConfig() (Config, error) {
	file, err := os.Open("config/config.json")
	if err != nil {
		return Config{}, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&GlobalConfig)
	return GlobalConfig, err
}

package infrastructure

import (
	"log"

	"github.com/spf13/viper"
)

var (
	cfg      *viper.Viper
	isLoaded bool = false
)

func _loadConfigFile() {
	log.Println("[INFO] Loading configuration file.")
	cfg = viper.New()
	cfg.AddConfigPath("./config")
	cfg.SetConfigName("app")
	cfg.SetConfigType("json")
}

// GetConfig ...
func GetConfig(key string) string {
	if isLoaded {
		return cfg.Get(key).(string)
	}

	_loadConfigFile()

	err := cfg.ReadInConfig()
	if err != nil {
		log.Printf("[ERROR] Fail to read configuration file: %s \n", err.Error())
		return ""
	}

	isLoaded = true
	return cfg.Get(key).(string)
}

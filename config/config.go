package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"reforce.pattern/pkg/controllers"
)

// Config структура конфига для сервиса prmObjects
type Config struct {
	AppVersion string `yaml:"app_version"`
	Server     struct {
		Port        string `yaml:"port"`
		Development bool   `yaml:"development"`
	} `yaml:"server"`
	Logger struct {
		DisableCaller     bool   `yaml:"disable_caller"`
		DisableStacktrace bool   `yaml:"disable_stacktrace"`
		Encoding          string `yaml:"encoding"`
		Level             string `yaml:"level"`
	} `yaml:"logger"`
	MongoDB struct {
		User       string   `yaml:"user"`
		Password   string   `yaml:"password"`
		DB         string   `yaml:"db"`
		Hosts      []string `yaml:"hosts"`
		Replica    string   `yaml:"replica"`
		App        string   `yaml:"app"`
		PoolLimits struct {
			Min uint64 `yaml:"min"`
			Max uint64 `yaml:"max"`
		} `yaml:"pool_limits"`
	} `yaml:"mongo_db"`
	ReforceID struct {
		URL       string `yaml:"url"`
		DecodeKey []byte `yaml:"decode_key"`
		ServiceID string `yaml:"service_id"`
	} `yaml:"reforce_id"`
	Cache struct {
		CleanupTime int64 `yaml:"cleanup_time"`
	} `yaml:"cache"`
	Cookie struct {
		Domain string `yaml:"domain"`
		MaxAge int64  `yaml:"max_age"`
	} `yaml:"cookies"`
}

var CFG = func() Config {
	var cfg Config
	cfg.openConfig()
	return cfg
}()

// ParseConfig Parse config file
func (config *Config) openConfig() {
	yamFile, err := os.Open(controllers.Path.ToFile("config/config.yaml"))
	if err != nil {
		log.Print("unable to decode into struct, %v", err)
		os.Exit(0)
	}

	defer func(f *os.File) {
		if err = f.Close(); err != nil {
			fmt.Printf("Error closing config file: %s\n", err)
		}
	}(yamFile)

	decoder := yaml.NewDecoder(yamFile)
	if err = decoder.Decode(config); err != nil {
		fmt.Println("Error decoding config file: " + err.Error())
		os.Exit(0)
	}

}

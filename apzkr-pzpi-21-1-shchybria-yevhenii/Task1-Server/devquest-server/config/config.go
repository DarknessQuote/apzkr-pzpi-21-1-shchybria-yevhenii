package config

import (
	"sync"

	"github.com/spf13/viper"
)

type (
	Server struct {
		Port int
	}

	Database struct {
		Host string
		Port int
		User string
		Password string
		DBName string
		SSLMode string
		TimeZone string
		ConnectTimeout int
	}

	Auth struct {
		Issuer string
		Secret string
		Audience string
		TokenExpiry int
		RefreshExpiry int
		CookieDomain string
		CookiePath string
		CookieName string
	}

	Config struct {
		Server *Server
		Database *Database
		Auth *Auth
	}
)

var (
	once sync.Once
	configInstance *Config
	configError error
)

func GetConfig() (*Config, error) {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./")

		if err := viper.ReadInConfig(); err != nil {
			configError = err
			return
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			configError = err
			return
		}
	})

	return configInstance, configError
}
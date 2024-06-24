package management

import (
	"sync"
	"time"

	"github.com/spf13/viper"
)

type (
	DeviceSettings struct {
		ID string
		Type string
	}

	ConnectionSettings struct {
		ServerHost string
		RequestInterval time.Duration
		ConnTimeout time.Duration
		MaxIdleConns int
		IdleConnTimeout time.Duration
		DisableCompression bool
	}

	DeviceConfig struct {
		DeviceSettings *DeviceSettings
		ConnectionSettings *ConnectionSettings
		UserID string
	}
)

var (
	once sync.Once
	configInstance *DeviceConfig
	configError error
)

func GetConfig() (*DeviceConfig, error) {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./")
		viper.SetDefault("userID", "")

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

func (c *DeviceConfig) SetConfigValue(key string, value any) error {
	viper.Set(key, value)
	viper.WriteConfig()

	if err := viper.Unmarshal(&configInstance); err != nil {
		return err
	}

	return nil
}
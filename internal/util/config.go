package util

import (
	"github.com/spf13/viper"
	"strings"
)

// Configurations exported
type Configurations struct {
	Server   ServerConfigurations
	Database DatabaseConfigurations
}

// ServerConfigurations exported
type ServerConfigurations struct {
	Http int
	Grpc int
}

// DatabaseConfigurations exported
type DatabaseConfigurations struct {
	DBName     string
	DBUser     string
	DBPassword string
}

func LoadConfig(path string) (config Configurations, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.SetEnvPrefix("APP")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return Configurations{}, err
	}

	err = viper.Unmarshal(&config)
	return config, nil
}

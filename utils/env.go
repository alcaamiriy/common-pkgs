package utils

import (
	"github.com/spf13/viper"
	"log"
)

var EnvConfig Config

type Config struct {
	LocalServerPort string `mapstructure:"LOCAL_SERVER_PORT"`
	SecretKey       string `mapstructure:"SECRET_KEY"`
	Environment     string `mapstruct:"environment"`
	LogLevel        string `mapstruct:"LOG_LEVEL"`
}

func LoadConfig(path string) (err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("failed to read env %v", err)
		return
	}

	err = viper.Unmarshal(&EnvConfig)
	if err != nil {
		log.Fatal("failed to unmarshal the viper env %v", err)
	}

	Logger.Info().Msg("Env file loaded successfully")
	return
}

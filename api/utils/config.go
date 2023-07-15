package utils

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func LoadConfig() error {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatal("Error trying to load config")
		return err
	}
	return nil
}

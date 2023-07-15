package main

import (
	"os"

	"github.com/brandon-charest/Shortify.git/handlers"
	"github.com/brandon-charest/Shortify.git/stores/redis"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		DisableColors: false,
	})
	logrus.SetOutput(os.Stdout)

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatalf("Error trying to load config: %v", err)
	}
	logrus.Info("Shortify App Start")
	initApp()

	logrus.Println("Shutting down...")
}

func initApp() error {

	_, err := redis.New()
	if err != nil {
		logrus.Fatalf("Could not setup redis: %v", err)
	}
	h, err := handlers.New()
	if err != nil {
		logrus.Fatalf("Could not setup app: %v", err)
	}
	logrus.Info("Connected to redis")

	if err := h.Listen(); err != nil {
		logrus.Fatalf("could not listen to http handlers: %v", err)
	}
	return nil
}

package main

import (
	"os"

	"github.com/brandon-charest/Shortify.git/api/handlers"
	"github.com/sirupsen/logrus"
)

func main() {

	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		DisableColors: false,
	})
	logrus.SetOutput(os.Stdout)

	initApp()

	logrus.Println("Shutting down...")
}

func initApp() error {
	h, err := handlers.New()
	if err != nil {
		logrus.Fatalf("Could not setup app: %v", err)
	}
	if err := h.Listen(); err != nil {
		logrus.Fatalf("could not listen to http handlers: %v", err)
	}
	return nil
}

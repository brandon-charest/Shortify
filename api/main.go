package main

import (
	"errors"
	"os"
	"os/signal"

	"github.com/brandon-charest/Shortify.git/api/handlers"
	"github.com/sirupsen/logrus"
)

func main() {
	kill := make(chan os.Signal, 1)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		DisableColors: false,
	})
	logrus.SetOutput(os.Stdout)
	signal.Notify(kill, os.Interrupt)

	<-kill
	logrus.Println("Shutting down...")
}

func init() error {
	h, err := handlers.New()

	if err != nil {
		return nil, errors.New("Could not setup app")
	}

	go func() {
		if err := h.listen(); err != nil {
			logrus.Fatalf("could not listen to http handlers: %v", err)
		}
	}()

}

package redis

import (
	"errors"
	"time"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

type Client struct {
	client *redis.Client
}

// func New() (*Client, error) {

// }

func newClient(hostaddr, password string, db int, maxRetries int, readTimeout string, writeTimeout string) (*Client, error) {
	var rt, wt time.Duration
	var err error

	if rt, err = time.ParseDuration(readTimeout); err != nil {
		logrus.Error(err)
		return nil, errors.New("Error parsing read timeout")
	}

	if wt, err = time.ParseDuration(writeTimeout); err != nil {
		logrus.Error(err)
		return nil, errors.New("Error parsing write timeout")
	}

	c := redis.NewClient(&redis.Options{
		Addr:         hostaddr,
		Password:     password,
		DB:           db,
		MaxRetries:   maxRetries,
		ReadTimeout:  rt,
		WriteTimeout: wt,
	})

	if _, err = c.Ping().Result(); err != nil {
		logrus.Error(err)
		return nil, errors.New("Error connecting to redis")
	}
	ret := &Client{client: c}
	return ret, nil
}

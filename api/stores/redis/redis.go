package redis

import (
	"errors"
	"time"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Client struct {
	client *redis.Client
}

func New() (*Client, error) {
	c, err := newClient(viper.GetString("DB_ADDRESS"),
		viper.GetString("DB_PASS"),
		viper.GetInt("DB_ID"),
		viper.GetInt("DB_MAX_RETY"),
		viper.GetString("DB_READ_TIMEOUT"),
		viper.GetString("DB_WRITE_TIMEOUT"))

	return c, err
}

func newClient(hostaddr, password string, db int, maxRetries int, readTimeout string, writeTimeout string) (*Client, error) {
	var rt, wt time.Duration
	var err error

	if rt, err = time.ParseDuration(readTimeout); err != nil {
		logrus.Error(err)
		return nil, errors.New("error parsing read timeout")
	}

	if wt, err = time.ParseDuration(writeTimeout); err != nil {
		logrus.Error(err)
		return nil, errors.New("error parsing write timeout")
	}

	c := redis.NewClient(&redis.Options{
		Addr:         hostaddr,
		Password:     password,
		DB:           db,
		MaxRetries:   maxRetries,
		ReadTimeout:  rt,
		WriteTimeout: wt,
	})
	logrus.Info("redis client created")
	if _, err = c.Ping().Result(); err != nil {
		logrus.Error(err)
		return nil, errors.New("error connecting to redis")
	}
	ret := &Client{client: c}
	return ret, nil
}

func (r *Client) Close() error {
	err := r.client.Close()
	return err
}

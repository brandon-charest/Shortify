package handlers

import (
	"errors"

	"github.com/brandon-charest/Shortify.git/stores/redis"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Handler struct {
	engine *gin.Engine
	store  redis.Client
}

func New(client redis.Client) (*Handler, error) {
	h := &Handler{engine: gin.New(), store: client}
	if err := h.setHandlers(); err != nil {
		return nil, errors.New("could not set handler")
	}

	return h, nil
}

func (h *Handler) setHandlers() error {
	h.engine.GET("/", h.resolveRoot)
	h.engine.POST("/shorten", h.resolveShorten)
	return nil
}

func (h *Handler) Listen() error {
	return h.engine.Run(viper.GetString("APP_PORT"))
}

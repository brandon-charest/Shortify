package handlers

import (
	"errors"

	"github.com/brandon-charest/Shortify.git/api/stores/redis"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	engine *gin.Engine
	store  redis.Client
}

func New() (*Handler, error) {
	h := &Handler{engine: gin.New()}
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
	return h.engine.Run(":3000")
}

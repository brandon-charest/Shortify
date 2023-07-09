package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/brandon-charest/Shortify.git/api/common"
	"github.com/gin-gonic/gin"
)

type request struct {
	URL       string        `json:"url" binding:"required"`
	ExpiresOn time.Duration `json:"expires_on"`
}

type response struct {
	URL            string        `json:"url" binding:"required"`
	ShortURL       string        `json:"short_url"`
	ExpiresOn      time.Duration `json:"expires_on"`
	XRateRemaining int
	XRateLimitRest time.Duration
}

type Handler struct {
	engine *gin.Engine
}

func New() (*Handler, error) {
	h := &Handler{engine: gin.New()}
	if err := h.setHandlers(); err != nil {
		return nil, errors.New("Could not set handler")
	}
	return h, nil
}

func (h *Handler) setHandlers() error {
	h.engine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusAccepted, gin.H{"message": "Shortify"})
	})

	return nil
}

func (h *Handler) listen() error {
	return h.engine.Run(common.APP_PORT)
}

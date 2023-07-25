package handlers

import (
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/brandon-charest/Shortify.git/stores/shared"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type request struct {
	URL string `json:"url" binding:"required"`
}

type response struct {
	URL            string        `json:"url" binding:"required"`
	ShortURL       string        `json:"short_url"`
	ExpiresOn      time.Duration `json:"expires_on"`
	XRateRemaining int
	XRateLimitRest time.Duration
}

func (h *Handler) resolveRoot(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, gin.H{"message": "Shortify"})
}

func (h *Handler) resolveShorten(ctx *gin.Context) {
	var data request
	if err := ctx.ShouldBind(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logrus.Debug("resolve shorten")

	if !govalidator.IsURL(data.URL) {
		ctx.JSON(http.StatusBadRequest, gin.H{"invalid URL": data.URL})
		return
	}
	id, err := h.store.CreateEntry(shared.Entry{
		URL: data.URL,
	})
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"redis error": err})
	}
	resp := response{
		URL:      data.URL,
		ShortURL: viper.GetString("DOMAIN") + "/" + id,
	}
	ctx.JSON(http.StatusAccepted, resp)
}

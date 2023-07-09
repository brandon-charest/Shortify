package handlers

import (
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

	ctx.JSON(http.StatusAccepted, gin.H{"message": "Shorten"})
}

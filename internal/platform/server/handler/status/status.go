package status

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// StatusHandler represents the status handler.

func StatusHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	}
}

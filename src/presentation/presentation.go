package presentation

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Route(BASE_URL string) error {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200, time.Now().String())
	})

	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, "status: ok")
	})

	return router.Run(BASE_URL)
}

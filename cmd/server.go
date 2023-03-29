package cmd

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func WebServer() {
	router := gin.Default()
	router.GET("/api/v1/question", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")

	})
	router.Run(":" + viper.GetString("server.port"))
}

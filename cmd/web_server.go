package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lightsoft/interview-knowledge-base/api"
	"github.com/spf13/viper"
)

func WebServer() {

	router := gin.Default()
	// v1

	v1 := router.Group("/api/v1")
	{
		v1.GET("/questions", api.QueryQuestion)
		v1.GET("/questions/:id", api.SearchQuestion)
		v1.POST("/questions/add", api.CreateQuestion)
		v1.PUT("/questions", api.UpdateQuestion)
		v1.DELETE("/questions", api.DeleteQuestion)
	}

	// Simple group: v2
	v2 := router.Group("/api/v2")
	{
		v2.GET("/questions", api.QueryQuestion)
		v2.GET("/questions/:id", api.SearchQuestion)
		v2.POST("/questions", api.CreateQuestion)
		v2.PUT("/questions", api.UpdateQuestion)
		v2.DELETE("/questions", api.DeleteQuestion)
	}

	srv := &http.Server{
		Addr:         ":" + viper.GetString("server.port"),
		Handler:      router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	go func() {
		// 监听请求
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 优雅Shutdown（或重启）服务
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt) // syscall.SIGKILL
	<-quit
	fmt.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
	}
	fmt.Println("Server exiting")

	// router.Run(":" + viper.GetString("server.port"))
}
func question(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

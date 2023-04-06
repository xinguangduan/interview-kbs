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
	"github.com/lightsoft/interview-knowledge-base/router"
	"github.com/spf13/viper"
)

func StartWebServer() {

	r := gin.Default()

	// router group
	router.ConfigRouter(r)

	srv := &http.Server{
		Addr:         ":" + viper.GetString("server.port"),
		Handler:      r,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	fmt.Printf("start server, %v, %v, %v", srv.Addr, srv.ReadTimeout, srv.WriteTimeout)

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

	<-ctx.Done()
	fmt.Println("Server exiting")

	// router.Run(":" + viper.GetString("server.port"))
}

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
	"github.com/lightsoft/interview-knowledge-base/configuration"
	"github.com/lightsoft/interview-knowledge-base/global"
	"github.com/lightsoft/interview-knowledge-base/middleware"
	"github.com/lightsoft/interview-knowledge-base/router"
	"github.com/lightsoft/interview-knowledge-base/utils"
	"github.com/spf13/viper"
)

func StartWebServer() {
	var initErr error
	r := gin.Default()

	// ===============================================================================
	// = 初始化日志组件
	global.Logger = configuration.InitLogger()
	// ===============================================================================
	// = 初始化Redis连接
	rdClient, err := configuration.InitRedis()
	global.RedisClient = rdClient
	if err != nil {
		initErr = utils.AppendError(initErr, err)
	}

	// ===============================================================================
	// = 初始化过程中, 遇到错误的最终处理
	if initErr != nil {
		if global.Logger != nil {
			global.Logger.Error(initErr.Error())
		}
		panic(initErr.Error())
	}

	r.Use(middleware.Cors())
	// router group
	router.ConfigRouter(r)

	srv := &http.Server{
		Addr:         ":" + viper.GetString("server.port"),
		Handler:      r,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	fmt.Printf("server started, %v, %v, %v ,version v1.0 \n", srv.Addr, srv.ReadTimeout, srv.WriteTimeout)

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

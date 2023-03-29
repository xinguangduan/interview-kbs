package cmd

import (
	"fmt"

	"github.com/lightsoft/interview-knowledge-base/configuration"
	"github.com/spf13/viper"
)

func Start() {
	configuration.Init()
	port := viper.Get("server.port")
	fmt.Println("start server ", port)
	WebServer()
}

func Clean() {
	fmt.Println("clean the cache data")
}

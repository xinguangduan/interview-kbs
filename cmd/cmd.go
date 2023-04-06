package cmd

import (
	"fmt"

	"github.com/lightsoft/interview-knowledge-base/configuration"
	"github.com/lightsoft/interview-knowledge-base/dao"
)

func Start() {
	configuration.Init()
	dao.InitDatabase()
	StartWebServer()
}

func Clean() {
	fmt.Println("clean the cache data")
}

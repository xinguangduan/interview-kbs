package main

import (
	"fmt"

	"github.com/lightsoft/interview-knowledge-base/cmd"
)

func main() {
	fmt.Println("start")
	defer cmd.Clean()

	cmd.Start()
}

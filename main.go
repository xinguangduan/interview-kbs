package main

import (
	"github.com/lightsoft/interview-knowledge-base/cmd"
)

func main() {
	defer cmd.Clean()

	cmd.Start()
}

func CheckError(err error) {
	panic("unimplemented")
}

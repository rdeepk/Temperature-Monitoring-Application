package main

import (
	"fmt"
	"github.com/rdeepk/temprature-monitoring/distributed/coordinator"
)

func main() {
	ql := coordinator.NewQueueListener()
	go ql.ListenForNewSource()

	var a string
	fmt.Scanln(&a)
}

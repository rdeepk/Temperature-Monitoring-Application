package main

import (
	"fmt"
	"github.com/rdeepk/temprature-monitoring/distributed/coordinator"
)

var dc *coordinator.DatabaseConsumer

func main() {
	ea := coordinator.NewEventAggregator()
	dc = coordinator.NewDatbaseConsumer(ea)
	ql := coordinator.NewQueueListener(ea)
	go ql.ListenForNewSource()

	var a string
	fmt.Scanln(&a)
}

package main

import (
	"github.com/carrot-go/rabbitagent"
	"os"
	"log"
	"context"
	"fmt"
	"time"
)

func main() {

	rabbitLoop, err := rabbitagent.NewRabbitLoop(
		"guest", "guest", "0.0.0.0:15672",
		&rabbitagent.Config{SysName:"test"});
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	metricLoop := rabbitagent.NewMetriceLoop(func(ctx context.Context, record *rabbitagent.MetricEntry) {
		fmt.Println("key:", record.Item, "| value:", record.Step)
	})

	metricC := make(chan []rabbitagent.MetricEntry, 1000)

	ticker := time.NewTicker(10 * time.Second)

	go metricLoop.Loop(metricC)
	go rabbitLoop.Loop(ticker.C, metricC)

	<- make(chan struct{})

	ticker.Stop()
	close(metricC)
}

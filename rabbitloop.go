package rabbitagent

import (
	"github.com/carrot-go/rabbitclient"
	"time"
	"fmt"
	"context"
)

type RabbitLoop struct {
	c    *rabbitclient.Conn
	conf *Config
}

func NewRabbitLoop(user, pwd, host string, conf *Config) (*RabbitLoop, error) {
	loop := &RabbitLoop{
		c: rabbitclient.NewConn(user, pwd, host),
		conf: conf,
	}
	return loop, nil
}

func (r *RabbitLoop) Loop(tick <-chan time.Time, metricC chan<- []MetricEntry) {
	mQueueC := make(chan []rabbitclient.Queue)
	mErrorC := make(chan error)
	for {
		select {
		case t := <-tick:
			ctx := context.TODO()
			fmt.Printf("rabbitLoop start at %s", t.String())
			go r.c.GetQueues(ctx, mQueueC, mErrorC)
			break
		case queues := <-mQueueC:
			metricC <- r.CollectQueueMetrics(queues)
			break
		case err := <-mErrorC:
			// TODO:
			panic(err)
			break
		}
	}
}

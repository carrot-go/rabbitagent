package rabbitagent

import (
	"context"
	"fmt"
)

type MetricLoop struct {
	recordEvent MetricFunc
}

func NewMetriceLoop(recordEvent MetricFunc) *MetricLoop {
	return &MetricLoop{recordEvent:recordEvent}
}

func (m *MetricLoop) Loop(metricC <-chan []MetricEntry) {
	for {
		for metrics := range metricC {
			fmt.Println("metric got..", len(metrics))
			for _, metric := range metrics {
				ctx := context.TODO()
				m.recordEvent(ctx, &metric)
			}
		}

	}
}

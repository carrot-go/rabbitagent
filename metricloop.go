package rabbitagent

import "context"

type MetricLoop struct {
	recordEvent MetricFunc
}

func NewMetriceLoop(recordEvent MetricFunc) *MetricLoop {
	return &MetricLoop{recordEvent:recordEvent}
}

func (m *MetricLoop) Loop(metricC <-chan []MetricEntry) {
	for {
		for metrics := range metricC {
			for _, metric := range metrics {
				ctx := context.TODO()
				m.recordEvent(ctx, &metric)
			}
		}

	}
}

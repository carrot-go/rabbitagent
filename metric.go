package rabbitagent

import "context"

type MetricEntry struct {
	SysName string
	Type    string
	Item    string
	Step    int64
}

type MetricFunc func(ctx context.Context, record *MetricEntry)

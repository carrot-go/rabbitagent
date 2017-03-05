package rabbitagent

import (
	"github.com/carrot-go/rabbitclient"
	"fmt"
)

func (r *RabbitLoop) CollectQueueMetrics(queues []rabbitclient.Queue) []MetricEntry {
	if len(queues) == 0 {
		return []MetricEntry{}
	}
	metrics := make([]MetricEntry, len(queues)*3)
	for i, queue := range queues {
		//
		metrics[i].SysName = r.conf.SysName
		metrics[i].Type = "biz"
		metrics[i].Item = fmt.Sprintf("queue.%s.depath", queue.Name)
		metrics[i].Step = queue.Messages

		offset := len(queues)
		metrics[i+offset].SysName = r.conf.SysName
		metrics[i+offset].Type = "biz"
		metrics[i+offset].Item = fmt.Sprintf("queue.%s.unacknowledged", queue.Name)
		metrics[i+offset].Step = queue.MessagesUnacknowledged

		offset = offset * 2
		metrics[i+offset].SysName = r.conf.SysName
		metrics[i+offset].Type = "biz"
		metrics[i+offset].Item = fmt.Sprintf("queue.%s.consumers", queue.Name)
		metrics[i+offset].Step = int64(queue.Consumers)
	}
	return metrics
}

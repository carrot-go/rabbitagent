package rabbitagent

import (
	"github.com/carrot-go/rabbitclient"
	"fmt"
)

func (r *RabbitLoop) CollectQueueMetrics(queues []rabbitclient.Queue) []MetricEntry {
	if len(queues) == 0 {
		return []MetricEntry{}
	}
	metrics := make([]MetricEntry, len(queues) * 2)
	for i, queue := range queues {
		//
		metrics[i].SysName = r.conf.SysName
		metrics[i].Type = "biz"
		metrics[i].Item = fmt.Sprintf("queue.%s.depath", queue.Name)
		metrics[i].Step = queue.Messages

		metrics[i].SysName = r.conf.SysName
		metrics[i].Type = "biz"
		metrics[i].Item = fmt.Sprintf("queue.%s.unacknowledged", queue.Name)
		metrics[i].Step = queue.MessagesUnacknowledged

		metrics[i].SysName = r.conf.SysName
		metrics[i].Type = "biz"
		metrics[i].Item = fmt.Sprintf("queue.%s.consumers", queue.Name)
		metrics[i].Step = int64(queue.Consumers)
	}
	return nil
}

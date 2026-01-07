package mqcmp

import (
	"fmt"
	"sync"

	"github.com/kweaver-ai/agent-go-common-pkg/src/infra/cmp/icmp"
	"github.com/kweaver-ai/agent-go-common-pkg/src/infra/common/chelper/cenvhelper"
	msqclient "devops.aishu.cn/AISHUDevOps/ONE-Architecture/_git/proton-mq-go"
)

var (
	mqOnce     sync.Once
	_mqClient  icmp.IMQClient
	configPath = "/sysvol/conf/mq/mq_config.yaml"
)

type mqClient struct {
	client                   msqclient.ProtonMQClient
	pollIntervalMilliseconds int64
	maxInFlight              int
}

// NewMQClientWithPath 根据路径创建消息队列
func NewMQClientWithPath(cfgPath ...string) icmp.IMQClient {
	if cenvhelper.IsLocalDev(cenvhelper.RunScenario_Aaron_Local_Dev) {
		return &mqClient{}
	}

	mqOnce.Do(func() {
		if len(cfgPath) != 0 {
			configPath = cfgPath[0]
		}

		mqSDK, err := msqclient.NewProtonMQClientFromFile(configPath)
		if err != nil {
			panic(fmt.Sprintf("[NewMQClientWithPath] ERROR: new mq client failed: %v\n", err))
		}

		_mqClient = &mqClient{
			client:                   mqSDK,
			pollIntervalMilliseconds: int64(100),
			maxInFlight:              16,
		}
	})

	return _mqClient
}

// Publish mq生产者
func (m *mqClient) Publish(topic string, msg []byte) (err error) {
	err = m.client.Pub(topic, msg)
	return
}

// Subscribe mq消费者
func (m *mqClient) Subscribe(topic, channel string, cmd func([]byte) error) (err error) {
	err = m.client.Sub(topic, channel, cmd, m.pollIntervalMilliseconds, m.maxInFlight)

	return
}

// Close 用于 Pub 内部的生产者关闭回收
func (m *mqClient) Close() {
	m.client.Close()
}

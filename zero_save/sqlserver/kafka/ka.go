package kafka

import (
	"time"

	"github.com/Shopify/sarama"
)

var (
	address = []string{"localhost:9092"}
)

type IKafka interface {
	SendMessage(msg *sarama.ProducerMessage) (partition int32, offset int64, err error)
	SendMessages(msgs []*sarama.ProducerMessage) error
}

type Kafka struct {
	ka sarama.SyncProducer
}

func NewKafka() IKafka {
	// 配置
	config := sarama.NewConfig()
	// 属性设置
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second
	// 创建生成者
	p, err := sarama.NewSyncProducer(address, config)
	if err != nil {
		return nil
	}
	// 最后关闭生产者
	// defer p.Close()

	return &Kafka{
		ka: p,
	}
}

func (k *Kafka) SendMessage(msg *sarama.ProducerMessage) (partition int32, offset int64, err error) {

	return 0, 0, nil
}

func (k *Kafka) SendMessages(msgs []*sarama.ProducerMessage) error {

	return nil
}

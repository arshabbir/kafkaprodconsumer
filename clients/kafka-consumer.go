package clients

type kafkaConsumer struct {
}

type KafkaConsumer interface {
	ReceiveMessage() ([]byte, error)
}

func NewConsumer() KafkaConsumer {
	return &kafkaConsumer{}
}

func (c *kafkaConsumer) ReceiveMessage() ([]byte, error) {
	return nil, nil
}

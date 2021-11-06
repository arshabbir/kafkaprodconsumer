package clients

import (
	"context"
	"fmt"
	"strconv"

	"github.com/segmentio/kafka-go"
)

type kafkaProducer struct {
	w        *kafka.Writer
	msgCount int
	status   chan int
}

type KafkaProducer interface {
	SendMessage(ctx context.Context, msg []byte) error
}

func NewProducer(brokers []string, topic string) KafkaProducer {
	//Connect to Kafka cluster a producer

	// intialize the writer with the broker addresses, and the topic
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: brokers,
		Topic:   topic,
	})

	return &kafkaProducer{w: w}
}

func (p *kafkaProducer) SendMessage(ctx context.Context, msg []byte) error {

	// each kafka message has a key and value. The key is used
	// to decide which partition (and consequently, which broker)
	// the message gets published on
	err := p.w.WriteMessages(ctx, kafka.Message{
		Key: []byte(strconv.Itoa((p.msgCount))),
		// create an arbitrary message payload for the value
		Value: []byte("this is message" + strconv.Itoa((p.msgCount))),
	})
	if err != nil {
		fmt.Println("could not write message " + err.Error())
		return err
	}

	// log a confirmation once the message is written
	fmt.Println("writes:", p.msgCount)
	p.msgCount++

	return nil
}

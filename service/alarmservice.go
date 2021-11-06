//go:generate mockgen -source alarmservice.go -destination mocks/alarmservicemock.go -package alarmservicemock AlarmService
package service

import (
	"context"
	"fmt"
	"time"

	"github.com/arshabbir/kafkaprodconsumer/clients"
)

type alarm struct {
	Id       string
	Name     string
	Severity string
}

type alarmService struct {
	mb clients.KafkaProducer
}

type AlarmService interface {
	StartConsumer(status chan int)
	Produce()
	Check() int
}

func NewAlarmService() AlarmService {
	//Read the env varables and set the brokers
	brokers := []string{"192.168.1.10"}
	topic := "test-topic"

	mb := clients.NewProducer(brokers, topic)
	return &alarmService{mb: mb}
}

func (s *alarmService) Check() int {
	return 0
}
func (s *alarmService) StartConsumer(status chan int) {

	//cnt := 0

	//Implement consumer to take the message from que and process it
	status <- 1

	for {
		time.Sleep(time.Millisecond)
		// fmt.Println("Consumer ...", cnt)
		// cnt++

		// if cnt > 10 {
		// 	status <- 2
		// 	close(status)
		// 	return
		// }

	}

}

func (s *alarmService) Produce() {
	ctx := context.Background()
	fmt.Println("Producer Started.")
	cnt := 0
	for {

		msg := fmt.Sprintf("Message : %d", cnt)
		if err := s.mb.SendMessage(ctx, []byte(msg)); err != nil {
			//Need to return the error
			break
		}
		time.Sleep(time.Second)
		cnt++
	}
}

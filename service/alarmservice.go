//go:generate mockgen -source alarmservice.go -destination mocks/alarmservicemock.go -package alarmservicemock AlarmService
package service

import (
	"fmt"
	"time"
)

type alarm struct {
	Id       string
	Name     string
	Severity string
}

type alarmService struct {
}

type AlarmService interface {
	StartConsumer(status chan int)
	Check() int
}

func NewAlarmService() AlarmService {
	return &alarmService{}
}

func (s *alarmService) Check() int {
	return 0
}
func (s *alarmService) StartConsumer(status chan int) {

	cnt := 0
	status <- 1

	for {
		time.Sleep(time.Millisecond)
		fmt.Println("Consumer ...", cnt)
		cnt++

		if cnt > 10 {
			status <- 2
			close(status)
			return
		}

	}

}

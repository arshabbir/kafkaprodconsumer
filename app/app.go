package app

import (
	"fmt"
	"time"

	"github.com/arshabbir/kafkaprodconsumer/service"
)

type app struct {
	Service service.AlarmService
}
type App interface {
	StartApplication(*int)
}

func NewApp(service service.AlarmService) App {
	return &app{Service: service}
}

func (ap *app) StartApplication(appStatus *int) {

	status := make(chan int)

	//Start the consumer
	go ap.Service.StartConsumer(status)
	//go ap.Service.Produce()

	for {
		select {
		case v, ok := <-status:

			if !ok {
				return
			}
			if v == 1 {
				*appStatus = 1
				fmt.Println("Consumer Started successfully")

			}
			if v == 2 {
				*appStatus = 2
				//fmt.Println("Application Stopped successfully")
				return
			}
		default:
			time.Sleep(time.Millisecond)

		}
	}

}

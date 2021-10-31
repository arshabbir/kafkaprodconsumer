package app

import (
	"fmt"
	"sync"
	"testing"

	alarmservicemock "github.com/arshabbir/kafkaprodconsumer/service/mocks"
	"github.com/golang/mock/gomock"
)

func TestStartApplicationConsumerStarted(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := alarmservicemock.NewMockAlarmService(ctrl)

	var appStatus int
	//status := make(chan int)
	app := NewApp(mockService)
	//status <- 1
	wg := &sync.WaitGroup{}

	wg.Add(1)
	defer wg.Wait()

	mockService.EXPECT().StartConsumer(gomock.Any()).Do(func(status chan int) {
		fmt.Println("StartConsumer....... ", app)
		status <- 1
		close(status)

		wg.Done()
	})

	app.StartApplication(&appStatus)

	if appStatus != 1 {
		//t.Fatalf("Appstatus value to be expected as 1 ", nil)
		t.Fail()
	}
	return
}

func TestStartApplicationApplicationStopped(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := alarmservicemock.NewMockAlarmService(ctrl)

	var appStatus int
	//status := make(chan int)
	app := NewApp(mockService)
	//status <- 1
	wg := &sync.WaitGroup{}

	wg.Add(1)
	defer wg.Wait()

	mockService.EXPECT().StartConsumer(gomock.Any()).Do(func(status chan int) {
		fmt.Println("StartConsumer....... ", app)
		status <- 2
		close(status)

		wg.Done()
	})

	app.StartApplication(&appStatus)

	if appStatus != 2 {
		//t.Fatalf("Appstatus value to be expected as 1 ", nil)
		t.Fail()
	}
	return
}

package main

import (
	"fmt"

	"github.com/arshabbir/kafkaprodconsumer/app"
	"github.com/arshabbir/kafkaprodconsumer/service"
)

func main() {

	fmt.Println("Starting the appliction ")
	var appStatus int

	app := app.NewApp(service.NewAlarmService())

	app.StartApplication(&appStatus)

	fmt.Printf("Application Stopped Successfully")

}

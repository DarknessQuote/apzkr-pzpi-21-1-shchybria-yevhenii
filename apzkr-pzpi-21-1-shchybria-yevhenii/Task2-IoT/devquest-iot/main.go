package main

import (
	"devquest-iot/app"
	"devquest-iot/management"
	"log"
)

func main() {
	configInstance, err := management.GetConfig()
	if err != nil {
		log.Panicln(err)
		return
	}

	app := app.CreateAppInstance(configInstance)
	
	err = app.Start()
	if err != nil {
		log.Panicln(err)
		return
	}
}
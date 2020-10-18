package main

import (
	"github.com/ashishjuyal/banking/app"
	"github.com/ashishjuyal/banking/logger"
)

func main() {

	logger.Info("Starting the application...")
	app.Start()

}

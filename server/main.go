package main

import (
	"github.com/ramonfsk/ibank/server/app"
	"github.com/ramonfsk/ibank/server/logger"
)

func main() {
	logger.Info("Starting app..")
	app.Start()
}

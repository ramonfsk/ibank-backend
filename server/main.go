package main

import (
	"github.com/ramonfsk/ibank-backend/server/app"
	"github.com/ramonfsk/ibank-backend/server/logger"
)

func main() {
	logger.Info("Starting app..")
	app.Start()
}

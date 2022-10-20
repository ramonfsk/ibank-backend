package main

import (
	"github.ibm.com/rfnascimento/ibank/server/app"
	"github.ibm.com/rfnascimento/ibank/server/logger"
)

func main() {
	logger.Info("Starting app..")
	app.Start()
}

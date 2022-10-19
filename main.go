package main

import (
	"github.ibm.com/rfnascimento/ibank/app"
	"github.ibm.com/rfnascimento/ibank/logger"
)

func main() {
	logger.Info("Starting app..")
	app.Start()
}

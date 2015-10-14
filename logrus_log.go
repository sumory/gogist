package main

import (
	log "github.com/Sirupsen/logrus"
	"./logrustest"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
}

func main() {



	logrustest.LogTwo()

	log.Warn("warn")

	logrustest.LogOne()

	log.Info("the end!")
}
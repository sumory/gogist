package main

import (
	log "github.com/Sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
}

func main() {
	log.WithFields(log.Fields{
		"animal": "1",
	}).Info("A walrus appears")

	log.Error("this is error")
	log.WithFields(log.Fields{
		"a":"b",
	})

	log.Warn("warn")
}
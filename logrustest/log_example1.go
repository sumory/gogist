package logrustest

import (
	logrus "github.com/Sirupsen/logrus"
)


var logger *logrus.Entry

func init() {
	logInstance := logrus.New()
	logInstance.Level = logrus.DebugLevel
	logInstance.Formatter = &logrus.TextFormatter{
		ForceColors:false,
		DisableColors: true,
		TimestampFormat: "2006-01-02 15:04:05.00000",
	}
	logger = logInstance.WithFields(logrus.Fields{
		"logger": "logger_a",
	})
}

func LogOne() {
	logger.WithField("1", "a").WithField("2", "b").Warn("this is warn for log")
}

func LogTwo(){
	logger.WithField("a", "1").WithField("b", "2").Info("this is info for log")
}
package logrustest

import (
	logrus "github.com/Sirupsen/logrus"
)


var bLogger *logrus.Entry

func init() {
	logInstance := logrus.New()
	logInstance.Level = logrus.DebugLevel
	logInstance.Formatter = &logrus.JSONFormatter{}
	bLogger = logInstance.WithFields(logrus.Fields{
		"logger": "logger_b",
	})


}

func Log1() {
	bLogger.WithField("1", "a").WithField("2", "b").Warn("this is warn for log")
}

func Log2(){
	bLogger.WithField("a", "1").WithField("b", "2").Info("this is info for log")
}
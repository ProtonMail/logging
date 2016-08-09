package main

import (
	pmlogger "github.com/ProtonMail/logging"
	"github.com/Sirupsen/logrus"
	goerror "github.com/go-errors/errors"
)

var log = setupLogger()

var env = "production"

func init() {
	log.Info("Init called")
}

func setupLogger() *pmlogger.Logger {
	if env == "production" { //production
		var logger = pmlogger.GetLogger("test 1")
		logger.Config("", "test.txt", logrus.DebugLevel)
		return logger
	}
	return pmlogger.GetDefaultLogger() //development
}

func main() {

	defer func() {
		err := recover()
		if err != nil {
			log.WithFields(logrus.Fields{
				"omg":     true,
				"details": goerror.Wrap(err, 4).ErrorStack(),
				"number":  100,
			}).Error("The ice breaks!")

			log.WithFields(logrus.Fields{
				"omg":     true,
				"details": goerror.Wrap(err, 4).ErrorStack(),
				"number":  100,
			}).Fatal("The ice breaks!")
		}
	}()

	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"number": 8,
	}).Debug("Started observing beach")

	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(logrus.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(logrus.Fields{
		"temperature": -4,
	}).Debug("Temperature changes")

	//
	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"number": 8,
	}).Debug("Started observing beach")

	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(logrus.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(logrus.Fields{
		"temperature": -4,
	}).Debug("Temperature changes")

	pmlogger.PrintStatus()

	log.WithFields(logrus.Fields{
		"animal": "orca",
		"size":   9009,
	}).Panic("It's over 9000!")

}

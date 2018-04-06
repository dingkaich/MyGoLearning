package mylogrus

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)

}

func MylogrusMain() {
	ConfigLocalFilesystemLogger("/Users/dingkai/Golang/MyGoLearning/src", "testlogrus.log", time.Hour*24, time.Hour*2)

	//log.WithFields(log.Fields{
	//	"animal": "walrus",
	//	"size":   10,
	//}).Info("A group of walrus emerges from the ocean")
	//
	//log.WithFields(log.Fields{
	//	"omg":    true,
	//	"number": 122,
	//}).Warn("The group's number increased tremendously!")
	//
	//log.WithFields(log.Fields{
	//	"omg":    true,
	//	"number": 100,
	//}).Fatal("The ice breaks!")
	//
	//// A common pattern is to re-use fields between logging statements by re-using
	//// the logrus.Entry returned from WithFields()
	requestLogger := log.WithFields(log.Fields{"request_id": 1, "user_ip": 1})
	requestLogger.Info("something happened on that request")
	requestLogger.Warn("something not great happened")
}

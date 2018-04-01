package mylogrus

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Create a new instance of the logger. You can have any number of instances.
var mylog = logrus.New()

func MylogrusMain1() {
	// The API for setting attributes is a little different than the package level
	// exported logger. See Godoc.
	mylog.Out = os.Stdout

	// You could set this to any `io.Writer` such as a file
	// file, err := os.OpenFile("logrus.mylog", os.O_CREATE|os.O_WRONLY, 0666)
	// if err == nil {
	//  mylog.Out = file
	// } else {
	//  mylog.Info("Failed to mylog to file, using default stderr")
	// }

	mylog.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	// logrus.SetFormatter(formatter)
}

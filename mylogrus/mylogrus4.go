package mylogrus

import (
	"flag"
	"fmt"
	"os"

	joonix "github.com/joonix/log"
	log "github.com/sirupsen/logrus"
)

func Mylogrusmain4() {
	lvl := flag.String("level", log.DebugLevel.String(), "log level")
	flag.Parse()

	level, err := log.ParseLevel(*lvl)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	log.SetLevel(level)
	log.SetFormatter(&joonix.FluentdFormatter{})

	log.Debug("hello world!")
}

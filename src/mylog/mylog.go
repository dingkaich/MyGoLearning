package mylog

import (
	"time"

	"runtime"

	"strconv"

	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/sirupsen/logrus"
)

func MylogMain1() {
	//log.Println("mylog")
	//log.SetFlags(log.Lshortfile | log.LstdFlags)
	//log.Println("file")
	//f, _ := os.OpenFile("test.log", os.O_RDWR|os.O_APPEND, 0640)
	//
	//log.SetOutput(f)
	//log.Println("this is file")
	simplelog()

}

type fileline struct{}

func (f *fileline) String() string {
	_, file, line, ok := runtime.Caller(12)
	if ok {
		return file + ":" + strconv.Itoa(line)
	}
	return "unkonwfile"
}

func simplelog() {
	//get write
	rotatelog, _ := rotatelogs.New("testmylog1.log.%Y%m%d%H%M%S",
		rotatelogs.WithMaxAge(time.Second*20),
		rotatelogs.WithRotationTime(time.Second*2))
	logrus.SetOutput(rotatelog)
	reallog := logrus.WithFields(logrus.Fields{
		"file": new(fileline),
	})

	timer := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-timer.C:
			reallog.Println("time out")
			goto final
		default:
			reallog.Print("goodboy")
			reallog.Print("is me")
		}
	}
final:
	reallog.Println("finish")
	mail163()
	logrus.Println("good")

}

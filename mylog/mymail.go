package mylog

import (
	"time"

	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/zbindenren/logrus_mail"
)

func mail163() {
	logger := logrus.New()
	mailhook, err := logrus_mail.NewMailAuthHook("dingkaich testmail",
		"smtp.163.com", 25,
		"dingkaich@163.com",
		"dingkaich@163.com",
		"dingkaich", "Ding1126")
	if err != nil {
		logger.AddHook(mailhook)
	} else {
		fmt.Println(err)
	}
	var filename = "123.txt"
	contextLogger := logger.WithFields(logrus.Fields{
		"file":    filename,
		"content": "GG",
	})
	//设置时间戳和message
	contextLogger.Time = time.Now()
	contextLogger.Message = "这是一个hook发来的邮件"

	//使用Fire发送,包含时间戳，message
	err = mailhook.Fire(contextLogger)
	if err != nil {
		fmt.Println(err)
	}
}

func Mylogmail() {
	mail163()
	logrus.Println("goodboy")

}

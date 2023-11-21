package example

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func init() {
	_ = os.MkdirAll("logs/"+time.Now().Format("200601"), 0600)
	file, err := os.OpenFile("logs/"+time.Now().Format("200601")+"/"+time.Now().Format("20060102T15")+".log", os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC, 0600)
	if err != nil {
		panic(err)
	}
	logrus.SetOutput(file)
	logrus.SetFormatter(&logrus.TextFormatter{DisableQuote: true})
	logrus.SetLevel(logrus.DebugLevel)
}

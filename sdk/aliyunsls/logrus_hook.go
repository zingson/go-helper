package aliyunsls

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

// AddHook 阿里云SLS存储日志
func AddHook(topic string, client *Client) {
	logrus.SetFormatter(&logrus.JSONFormatter{}) // 只有JSON格式输出的日志才能写入阿里云SLS
	logrus.AddHook(&SlsHook{
		topic: topic,
		c:     client,
	})
	logrus.Info("Add SlsHook success ")
}

//SlsHook 日志Hook
type SlsHook struct {
	topic string
	c     *Client
}

func (hook *SlsHook) Fire(entry *logrus.Entry) (err error) {
	bytes, err := entry.Bytes()
	if err != nil {
		return
	}
	var mkv map[string]string // 转换为Map，存入SLS
	err = json.Unmarshal(bytes, &mkv)
	if err != nil {
		return
	}
	hook.c.Log(hook.topic, mkv)
	return
}

func (hook *SlsHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

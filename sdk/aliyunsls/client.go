package aliyunsls

import (
	"fmt"
	"github.com/aliyun/aliyun-log-go-sdk/producer"
	"time"
)

// Option 阿里云日志服务参数
type Option struct {
	Endpoint        string `json:"endpoint" toml:"endpoint"`
	AccessKeyID     string `json:"access_key_id"  toml:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret" toml:"access_key_secret"`
	Project         string `json:"project" toml:"project"`
	Logstore        string `json:"logstore" toml:"logstore"`
}

type Client struct {
	Option
	producer *producer.Producer
	source   string
}

func NewClient(option Option) *Client {
	return new(Client).SetOption(option)
}

func (c *Client) SetOption(option Option) *Client {
	c.Option = option
	if c.producer != nil {
		c.producer.SafeClose()
	}
	c.source = LocalIP()
	c.producer = c.Producer()
	return c
}

func (c *Client) Producer() *producer.Producer {
	producerConfig := producer.GetDefaultProducerConfig()
	producerConfig.Endpoint = c.Endpoint
	producerConfig.AccessKeyID = c.AccessKeyID
	producerConfig.AccessKeySecret = c.AccessKeySecret
	c.producer = producer.InitProducer(producerConfig)
	c.producer.Start() // 启动producer实例
	return c.producer
}

//Log 记录日志  { msg:"" }
func (c *Client) Log(topic string, kv map[string]string) {
	err := c.producer.SendLog(c.Project, c.Logstore, topic, c.source, producer.GenerateLog(uint32(time.Now().Unix()), kv))
	if err != nil {
		fmt.Println("SLS SendLog Error:", err.Error())
	}
}

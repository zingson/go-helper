package hmongo

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/zingson/go-helper/hslice"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
	"sync"
	"time"
)

/**
Key-Value 集合,用于保存配置

// 新建Client
client = kv.NewKv(dsn,name)

// 读取配置
opt1 = hmongo.Get(client,key)
opt2 = hmongo.Get(client,key)
optN = hmongo.Get(client,key)

*/

// NewKv 客户端
// dsn Mongodb数据库连接uri
// name 集合名
// keys 读取的配置key名字数组，为空读取所有
func NewKv(dsn, name string) *Client {
	c := &Client{
		cache:      sync.Map{},
		collection: NowDatabase(dsn).Collection(name),
		keys:       []string{},
	}
	return c
}

type Client struct {
	cache      sync.Map // value为 []byte
	collection *mongo.Collection
	keys       []string
}

// Kv 配置集合
type Kv struct {
	Paas   string         `bson:"paas" json:"paas"`     // 表态标识
	Key    string         `bson:"key" json:"key"`       // 配置Key
	Value  map[string]any `bson:"value" json:"value"`   // 配置对象， 配置字段特殊key： key_label 显示名称 ,key_remark 备注说明
	Label  string         `bson:"label" json:"label"`   // 显示名称
	Remark string         `bson:"remark" json:"remark"` // 备注说明
}

// Get 读指定Key的配置
func (c *Client) Get(key string, opt any) {
	// 从缓存读取
	if c.cacheLoad(key, opt) {
		return
	}
	// 从表中读取
	c.Load(key)
	// 从缓存读取
	if c.cacheLoad(key, opt) {
		return
	}
	panic(errors.New(fmt.Sprintf("ERROR:缺少【%s】配置，请检查配置项", key)))
}

// Get 读指定Key的配置
func Get[T any](c *Client, key string) (v T) {
	c.Get(key, &v)
	return
}

// 从缓存读
func (c *Client) cacheLoad(key string, value any) bool {
	/*k := key + "_"
	if b, ok := c.cache.Load(k); ok {
		value = &b
		return true
	}*/

	if b, ok := c.cache.Load(key); ok {
		err := json.Unmarshal(b.([]byte), value)
		if err != nil {
			panic(fmt.Errorf("key=%s 配置解析异常 %s", key, err.Error()))
		}
		//c.cache.Store(k, value) // 缓存解析结果
		return true
	}
	return false
}

// Load 加载数据到缓存
func (c *Client) Load(keys ...string) {
	if len(keys) == 0 {
		return
	}
	list, err := DaoFind[Kv](nil, c.collection, bson.D{bson.E{"key", bson.M{"$in": keys}}})
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	for _, item := range list {
		var b []byte
		b, err = json.Marshal(item.Value)
		if err != nil {
			panic(fmt.Errorf("key=%s 配置非JSON格式 %s", item.Key, err.Error()))
		}
		c.cache.Store(item.Key, b)
		c.keys = hslice.AppendUnique[string](c.keys, item.Key)
	}
}

// Watch 配置更新
func (c *Client) Watch() {
	go func(c *Client) {
		for {
			// TODO 待完成，定时改为 Mongo Watch
			time.Sleep(5 * time.Minute) // 5分钟执行一次
			c.Load(c.keys...)
			logrus.Infof("Watch: %s", strings.Join(c.keys, ","))
		}
	}(c)
}

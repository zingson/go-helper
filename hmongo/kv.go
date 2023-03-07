package hmongo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
	"sync"
	"time"
)

/*******

Key-Value 集合,用于保存全局配置

方式1：加载KV，如果有更新，需重新调用Load函数
var kv = NewKv(dsn,name,keys).Load()

方式2：监听KV更新
var kv = NewKv(dsn,name,keys).Watch()

*/

//NewKV 新建配置对象
//dsn Mongodb数据库连接uri
//name 集合名
//keys 读取的配置key名字数组，为空读取所有
func NewKV(dsn, name string, keys ...string) *Kv {
	collection := NowDatabase(dsn).Collection(name)
	c := &Kv{
		cache:      sync.Map{},
		collection: collection,
		keys:       keys,
	}
	return c
}

type Kv struct {
	cache      sync.Map // value为 []byte
	collection *mongo.Collection
	keys       []string
}

// Get 读取指定配置
func (c *Kv) Get(key string, value interface{}) {
	if c == nil {
		panic("kv is nil")
	}
	v, ok := c.cache.Load(key)
	if !ok || v == nil {
		c.LoadItem(key)
		v, ok = c.cache.Load(key)
		if !ok || v == nil {
			panic(errors.New(fmt.Sprintf("ERROR:未读取到【%s】配置，请检查项目配置", key)))
		}
	}

	err := json.Unmarshal(v.([]byte), value)
	if err != nil {
		panic(errors.New(fmt.Sprintf("ERROR:无法解析【%s】配置，请检查类型是否正确", key)))
	}
	return
}

// GetCache 读取所有配置项
func (c *Kv) GetCache() map[string]string {
	m := make(map[string]string)
	c.cache.Range(func(key, value interface{}) bool {
		m[key.(string)] = string(value.([]byte))
		return true
	})
	return m
}

// GetItem 读取其中一项配置
func (c *Kv) GetItem(name string) (v []byte) {
	value, ok := c.cache.Load(name)
	if ok {
		v = value.([]byte)
	}
	return
}

// Load 加载配置
func (c *Kv) Load() *Kv {
	filter := bson.M{}
	if c.keys != nil && len(c.keys) > 0 {
		bsonA := bson.A{}
		for _, key := range c.keys {
			bsonA = append(bsonA, key)
		}
		filter["key"] = bson.M{"$in": bsonA}
	}

	cur, err := c.collection.Find(nil, filter)
	if err != nil {
		panic(err)
	}
	var configs []*KvObject
	err = cur.All(nil, &configs)
	if err != nil {
		panic(err)
	}
	for _, conf := range configs {
		value, _ := json.Marshal(conf.Value)
		if value == nil {
			continue
		}
		c.cache.Store(conf.Key, value)
	}
	fmt.Println("Mongo KV Load finish. ")
	return c
}

func (c *Kv) LoadItem(key string) {
	var config *KvObject
	err := c.collection.FindOne(nil, bson.M{"key": key}).Decode(&config)
	if err != nil {
		logrus.Errorf("Mongo table=%s  key=%s Error:%s", c.collection.Name(), key, err.Error())
		panic(err)
	}
	value, _ := json.Marshal(config.Value)
	if value == nil {
		return
	}
	if !strings.Contains(strings.Join(c.keys, ","), key) {
		c.keys = append(c.keys, key)
	}
	c.cache.Store(config.Key, value)
}

// Watch 监听配置更新
func (c *Kv) Watch() *Kv {
	c.Load()
	go func(kv *Kv) {
		ctx := context.TODO()
		cs, err := kv.collection.Watch(ctx, mongo.Pipeline{})
		if err != nil {
			fmt.Println(err.Error())
			time.Sleep(5 * time.Minute)
			kv.Watch() // 异常情况，间隔5分钟重试
		}
		defer cs.Close(ctx)
		for cs.Next(ctx) {
			fmt.Println("Watch config ", cs.Current)
			kv.Load()
		}
	}(c)
	fmt.Println("Mongo KV Watch ... ")
	return c
}

type KvObject struct {
	Key    string                 `bson:"key" json:"key"`       // 配置Key
	Value  map[string]interface{} `bson:"value" json:"value"`   // 配置对象，键值均为字符串类型，key_label 标签名称 ,key_remark 备注说明
	Label  string                 `bson:"label" json:"label"`   // 标签名称
	Remark string                 `bson:"remark" json:"remark"` // 备注说明
}

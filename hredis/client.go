package hredis

import (
	"github.com/go-redis/redis/v8"
	"sync"
)

type Option struct {
	Dsn string `json:"dsn"`
}

// newClient
// redisURL无密码示例 redis://101.133.221.239:7777/0
// redisURL带密码示例 redis://user:password@localhost:6789/3?dial_timeout=3&db=1&read_timeout=6s&max_retries=2
func newClient(redisURL string) *redis.Client {
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		panic(err)
	}
	return redis.NewClient(opt)
}

var _cache sync.Map

// NowClient 当前客户端
func NowClient(url string) (client *redis.Client) {
	if v, ok := _cache.Load(url); ok {
		client = v.(*redis.Client)
		return
	}
	client = newClient(url)
	_cache.Store(url, client)
	return
}

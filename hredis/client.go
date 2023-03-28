package hredis

import (
	"github.com/redis/go-redis/v9"
	"strings"
	"sync"
)

var (
	//ClientName 连接名字，可通过 ` hredis.ClientName="Name" ` 改为项目名或者具体任务名称
	ClientName = "hredis"

	// Redis Client 缓存
	_cache sync.Map
)

// NowClient 获取缓存客户端，不存在则创建
func NowClient(redisURL string) (client redis.UniversalClient) {
	if v, ok := _cache.Load(redisURL); ok {
		client = v.(redis.UniversalClient)
		return
	}
	client = newClient(redisURL)
	_cache.Store(redisURL, client)
	return
}

// newClient
func newClient(redisURL string) redis.UniversalClient {
	universalOpt := &redis.UniversalOptions{
		Addrs:      []string{},
		ClientName: ClientName,
		DB:         0,
	}
	if strings.Contains(redisURL, "addr=") {
		//Redis Cluster 连接
		//redis://user:password@localhost:6789?dial_timeout=3&read_timeout=6s&addr=localhost:6790&addr=localhost:6791
		opt, err := redis.ParseClusterURL(redisURL)
		if err != nil {
			panic(err)
		}
		// 集群只能使用db0
		universalOpt.Addrs = opt.Addrs
	} else {
		//单实例连接
		//redis://user:password@localhost:6789/3?dial_timeout=3&db=1&read_timeout=6s&max_retries=2
		opt, err := redis.ParseURL(redisURL)
		if err != nil {
			panic(err)
		}
		universalOpt.DB = opt.DB
		universalOpt.Addrs = []string{opt.Addr}
	}
	return redis.NewUniversalClient(universalOpt)
}

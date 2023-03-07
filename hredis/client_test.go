package hredis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"testing"
)

func TestNew(t *testing.T) {
	client := NowClient("redis://:Himkt2022@r-uf60buxhzyv6oilsdvpd.redis.rds.aliyuncs.com:6379/0")

	err := client.Set(context.Background(), "k2", "k2", -1).Err()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("set success")
}

func TestNew2(t *testing.T) {
	client := redis.NewClient(&redis.Options{Addr: "r-uf60buxhzyv6oilsdvpd.redis.rds.aliyuncs.com:6379", DB: 0})
	err := client.Set(context.Background(), "k2", "k2", -1).Err()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("set success")
}

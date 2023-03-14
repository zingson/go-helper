package test

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/redis/go-redis/v9"
	"testing"
	"time"
)

/* Redis Stream */

const (
	stream = "helper"
	group  = "group4"
)

// 生产者
func TestProducer(t *testing.T) {
	for i := 0; i < 10; i++ {
		err := client.XAdd(context.Background(), &redis.XAddArgs{
			Stream: stream,
			Values: map[string]string{"vid": fmt.Sprintf("1001_%d", time.Now().Unix())},
		}).Err()
		if err != nil {
			t.Error(err.Error())
		}
		time.Sleep(time.Second)
	}

}

// 查询队列数据，只读数据不删数据
func TestConsumerXRead(t *testing.T) {
	xStreamSlice := client.XRead(context.Background(), &redis.XReadArgs{
		Streams: []string{"helper", "0"},
		Count:   10,
	})
	err := xStreamSlice.Err()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("len(xStream)=%d", len(xStreamSlice.Val()))
	for _, stream := range xStreamSlice.Val() {
		for _, msg := range stream.Messages {
			t.Logf("ID:%s  values:%v", msg.ID, msg.Values)
		}
	}
}

// 消费组
func TestConsumer(t *testing.T) {

	for {
		client.XGroupCreateMkStream(context.Background(), stream, group, "0")
		xStreamSlice := client.XReadGroup(context.Background(), &redis.XReadGroupArgs{
			Group:    group,
			Consumer: "group3-1",
			Streams:  []string{stream, ">"},
			Count:    1,
			Block:    time.Minute,
			NoAck:    false,
		})
		err := xStreamSlice.Err()
		if err != nil {
			t.Error(err)
			continue
		}
		t.Logf("len(xStream)=%d", len(xStreamSlice.Val()))
		t.Logf("%v", xStreamSlice.Val())
		for _, xStream := range xStreamSlice.Val() {
			for _, msg := range xStream.Messages {
				t.Logf("ID:%s  values:%v", msg.ID, msg.Values)
				client.XAck(context.Background(), stream, group, msg.ID)
			}
		}
	}

}

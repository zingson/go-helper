package test

import (
	"github.com/zingson/goh/hcron"
	"github.com/zingson/goh/htime"
	"testing"
	"time"
)

// singleNode 存储实现示例
type SingleNode struct{}

func (*SingleNode) Find() []*hcron.Entry                { println("Find"); return make([]*hcron.Entry, 0) }
func (*SingleNode) Save(schedule *hcron.Entry)          { println(schedule.Key, "Save") }
func (*SingleNode) Lock(key string) (ok bool)           { println(key, "Lock"); return true }
func (*SingleNode) Unlock(key string)                   { println(key, "Unlock") }
func (*SingleNode) Log(schedule *hcron.Entry, ms int64) { println("Log ", schedule.Key, " ", ms, "ms") }

func TestA(t *testing.T) {
	a := hcron.New(new(SingleNode))
	a.AddFunc("AAA", "0/1 * * * * *", func() {
		t.Log("定时执行 ", htime.NowF19())
	})
	time.Sleep(50 * time.Minute)
}

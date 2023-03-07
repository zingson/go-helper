package hcron

import (
	"root/src/pkg/htime"
	"testing"
	"time"
)

// singleNode 存储实现示例
type singleNode struct{}

func (*singleNode) Find() []*Entry                { println("Find"); return make([]*Entry, 0) }
func (*singleNode) Save(schedule *Entry)          { println(schedule.Key, "Save") }
func (*singleNode) Lock(key string) (ok bool)     { println(key, "Lock"); return true }
func (*singleNode) Unlock(key string)             { println(key, "Unlock") }
func (*singleNode) Log(schedule *Entry, ms int64) { println("Log ", schedule.Key, " ", ms, "ms") }

func TestA(t *testing.T) {
	a := New(new(singleNode))
	a.AddFunc("AAA", "0/1 * * * * *", func() {
		t.Log("定时执行 ", htime.NowF19())
	})
	time.Sleep(50 * time.Minute)
}

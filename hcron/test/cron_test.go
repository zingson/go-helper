package test

import (
	"github.com/zingson/go-helper/hcron"
	"testing"
	"time"
)

// singleNode 存储实现示例
type SingleNode struct{}

func (*SingleNode) Find() []*hcron.Entry {
	println("Find")
	return make([]*hcron.Entry, 0)
}
func (*SingleNode) Save(entry *hcron.Entry) {
	println("Save", entry.Key)
}
func (*SingleNode) Lock(key string, p, n time.Time) (ok bool) {
	println("Lock", key, p.String(), n.String())
	return true
}
func (*SingleNode) Log(entry *hcron.Entry, ms int64) {
	println("Log ", entry.Key, " ", ms, "ms")
}

func TestA(t *testing.T) {
	/*a := hcron.New(new(SingleNode))
	a.AddFunc("AAA", "0/3 * * * * *", func() {
		t.Log("定时执行 ", htime.NowF19())
	})
	time.Sleep(50 * time.Minute)*/
}

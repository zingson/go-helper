package test

import (
	"github.com/zingson/go-helper/hcache"
	"github.com/zingson/go-helper/hid"
	"sync"
	"testing"
	"time"
)

type Response struct {
	Rid string `json:"rid"`
}

func TestMemory(t *testing.T) {

	res, err := hcache.Load[*Response](hcache.Memory, "kk", time.Minute, func() (out any, err error) {
		out = &Response{Rid: hid.UUID32()}
		return
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(res)

}

func TestMemory2(t *testing.T) {
	t.Log(time.Now().Format(time.RFC3339))
	var key string
	for i := 0; i < 1000000; i++ {
		key = hid.UUID32()
		_, err := hcache.Load[*Response](hcache.Memory, key, time.Minute, func() (out any, err error) {
			out = &Response{Rid: hid.UUID32()}
			return
		})
		if err != nil {
			t.Error(err)
			break
		}

	}
	t.Log(time.Now().Format(time.RFC3339))
	_, _ = hcache.Load[*Response](hcache.Memory, key, time.Minute, func() (out any, err error) {
		out = &Response{Rid: hid.UUID32()}
		return
	})
	t.Log(time.Now().Format(time.RFC3339))
}

func TestMemory3(t *testing.T) {
	t.Log(time.Now().Format(time.RFC3339))
	var m = sync.Map{}
	for i := 0; i < 1000000; i++ {
		key := hid.UUID32()
		m.Store(key, &Response{Rid: hid.UUID32()})
		_, _ = m.Load(key)
	}
	t.Log(time.Now().Format(time.RFC3339))
}

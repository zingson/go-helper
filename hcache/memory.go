package hcache

import (
	"sync"
	"time"
)

// Memory 使用内存Map作为应用缓存
var Memory = new(MemoryImpl)

type MemoryImpl struct {
	syncMap sync.Map
}

type value struct {
	v   any       // 缓存数据
	exp time.Time // 过期时间
}

// 定时删除过期key
func init() {
	go func() {
		for {
			time.Sleep(11 * time.Minute)
			Memory.clear()
		}
	}()
}

// 清理过期Key
func (m *MemoryImpl) clear() {
	m.syncMap.Range(func(key, val any) bool {
		v := val.(value)
		if time.Now().After(v.exp) {
			m.syncMap.Delete(key)
		}
		return true

	})
}

// Load 读取数据，存在则直接返回，不存在则保存缓存结果再返回
// key 缓存key必须全局唯一
// duration 缓存有效时间
// ds 当缓存没有数据时从数据源读取
func (m *MemoryImpl) Load(key string, duration time.Duration, ds func() (out any, err error)) (out any, err error) {
	// 读取缓存
	if valueAny, ok := m.syncMap.Load(key); ok {
		item := valueAny.(value)
		if time.Now().Before(item.exp) {
			out = item.v
			return
		}
	}
	// 读数据源
	out, err = ds()
	if err != nil || out == nil {
		return
	}
	// 写入缓存
	m.syncMap.Store(key, value{v: out, exp: time.Now().Add(duration)})
	return
}

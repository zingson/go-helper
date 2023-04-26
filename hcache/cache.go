package hcache

import "time"

type Cache interface {
	// 读取数据
	// key 缓存key必须全局唯一
	// duration 缓存有效时间
	// ds 当缓存没有数据时从数据源读取
	Load(key string, duration time.Duration, ds func() (out any, err error)) (out any, err error)
}

func Load[T any](c Cache, key string, duration time.Duration, ds func() (out any, err error)) (out T, err error) {
	r, err := c.Load(key, duration, ds)
	if err != nil {
		return
	}
	out = r.(T)
	return
}

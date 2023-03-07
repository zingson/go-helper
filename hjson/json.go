package hjson

import "encoding/json"

// Stringify 对象转为json字符串
func Stringify(v any) string {
	b, _ := json.Marshal(v)
	return string(b)
}

// Parse 字符串解析为JSON对象
func Parse[T any](s string) (r T, err error) {
	err = json.Unmarshal([]byte(s), &r)
	if err != nil {
		return
	}
	return
}

// Convert 对象通过json中转转换
func Convert[S any, T any](source S) (target T, err error) {
	b, _ := json.Marshal(source)
	err = json.Unmarshal(b, &target)
	if err != nil {
		return
	}
	return
}

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

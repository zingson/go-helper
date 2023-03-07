package hmongo

import (
	"fmt"
	"io/ioutil"
	"os"
)

func NewKvLoad(keys ...string) *Kv {
	return NewKV(DsnLoad(), KvName(), keys...).Load()
}

func NewKvWatch(keys ...string) *Kv {
	return NewKV(DsnLoad(), KvName(), keys...).Watch()
}

// DsnLoad 加载Dsn配置，默认方案
func DsnLoad(files ...string) string {

	//从环境变量读取
	connStr := os.Getenv("HMONGO")
	if connStr != "" {
		return connStr
	}

	// 从 docker secrets 读取
	filename := "/run/secrets/hmongo"
	if len(files) > 0 {
		filename = files[0]
	}
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error hmongo file:", err.Error())
		return ""
	}

	return string(bs)
}

// KvName 集合名称
func KvName() string {
	name := "dsys_config"
	defer func() {
		fmt.Println("KvName: " + name)
	}()
	if denv := os.Getenv("DENV"); denv == "local" {
		name = name + "_" + denv
		return name
	}
	if bs, _ := ioutil.ReadFile("/run/secrets/denv"); bs != nil && string(bs) == "local" {
		name = name + "_" + string(bs)
		return name
	}
	return name
}

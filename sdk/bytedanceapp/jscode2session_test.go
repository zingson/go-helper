package bytedanceapp

import (
	"errors"
	"github.com/BurntSushi/toml"
	"testing"
)

func tomlGet() (config *Config) {
	_, err := toml.DecodeFile("D:\\Projects\\hlib-go\\helper\\sdk-bytedanceapp\\.secret\\fxc.toml", &config)
	if err != nil {
		panic(err)
	}
	if config == nil {
		panic(errors.New("未读取到配置文件"))
	}
	return
}

func TestJscode2session(t *testing.T) {
	config := tomlGet()

	code := "6c60T_569d3aIccQeI6acvkr_se0x27c804of0JuWQ9BnU4qk3b9CxHWCs6sc4G3kaWWFwnJC--e4lchYTFaUjx3YeTz8_RnwzcHuNeAv0awFg4g_fhUou7NYhE"
	data, err := Jscode2session(config, code)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(data)
}

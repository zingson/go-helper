package hmongo

import (
	"os"
)

// Dsn 读取数据库连接配置
// 读取顺序
// 1.环境变量
// 2.配置文件
func Dsn(files ...string) (v string) {

	//从环境变量读取
	v = os.Getenv("HMONGO")
	if v != "" {
		return
	}

	// 从文件读取， 最后从 docker secrets 默认路径读取
	files = append(files, "/run/secrets/hmongo")
	for _, path := range files {
		b, err := os.ReadFile(path)
		if err != nil {
			continue
		}
		v = string(b)
		return
	}

	return
}

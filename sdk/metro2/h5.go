package metro2

import "fmt"

//H5Detail 乘车二维码页面
func H5Detail(config *Config, token, guid string) string {
	return fmt.Sprintf("%s?token=%s&guid=%s", config.H5url, token, guid)
}

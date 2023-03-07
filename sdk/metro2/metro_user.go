package metro2

import (
	"net/http"
)

//ResUserInfo 获取用户信息，给地铁查询用户信息
func ResUserInfo(config *Config, rid string, request *http.Request, response http.ResponseWriter, f func(AppToken string) (data *ResUserInfoData, err error)) {
	httpCallback(config, rid, request, response, func(reqBody string) (resData any, err error) {
		rmap, err := jsonParse[map[string]string](reqBody)
		if err != nil {
			return
		}
		resData, err = f(rmap["AppToken"])
		if err != nil {
			return
		}
		return
	})
}

type ResUserInfoData struct {
	GuidUser    string `json:"GuidUser"`
	MobilePhone string `json:"MobilePhone"`
}

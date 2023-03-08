package douyinpoi

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"
)

// 通用错误码
type errorCode int64

const (
	// 通用状态码
	ERROR_CODE_0       errorCode = 0       // 成功
	ERROR_CODE_2100004 errorCode = 2100004 // 系统繁忙，此时请开发者稍候再试
	ERROR_CODE_2100005 errorCode = 2100005 // 参数不合法
	ERROR_CODE_2100007 errorCode = 2100007 // 无权限操作
	ERROR_CODE_2100009 errorCode = 2100009 // 用户被禁封使用该操作
	ERROR_CODE_2190001 errorCode = 2190001 // quota已用完
	ERROR_CODE_2190004 errorCode = 2190004 // 应用未获得该能力
	ERROR_CODE_2190015 errorCode = 2190015 // 请求参数access_token openid不匹配
	ERROR_CODE_2190016 errorCode = 2190016 // 当前应用已被封禁或下线
)

const (
	TOKEN_GRAN_TYPE = "client_credential" //类别
)

//常用域名
const (
	WEBSITE_URl                             = "https://open.douyin.com/"                      // 域名
	REQUEST_ADDRESS                         = WEBSITE_URl + "oauth/client_token/"             // 获取token
	QUERY_POI_REQUEST_ADDRESS               = WEBSITE_URl + "poi/query/"                      // 获取抖音 POI ID
	SUPPLIER_MATACH_REQUEST_ADDRESS         = WEBSITE_URl + "poi/v2/supplier/match/"          // 发起店铺匹配POI同步任务
	SUPPLIER_QUERY_REQUEST_ADDRESS          = WEBSITE_URl + "poi/supplier/query/"             // 查询店铺
	SUPPLIER_QUERY_SUPPLIER_REQUEST_ADDRESS = WEBSITE_URl + "poi/v2/supplier/query/supplier/" // 店铺匹配状态查询
	SUPPLIER_QUERY_TASK_REQUEST_ADDRESS     = WEBSITE_URl + "poi/v2/supplier/query/task/"     // 店铺匹配任务结果查询
	SUPPLIER_SYNC_REQUEST_ADDRESS           = WEBSITE_URl + "poi/supplier/sync/"              // 商铺同步
	SPU_GET_REQUEST_ADDRESS                 = WEBSITE_URl + "poi/v2/spu/get/"                 // 多门店SPU信息查询
	SPU_STATUS_SYNC_REQUEST_ADDRESS         = WEBSITE_URl + "poi/v2/spu/status_sync/"         // 多门店SPU状态同步(主要商品下线)
	SPU_STOCK_UPDATE_REQUEST_ADDRESS        = WEBSITE_URl + "poi/v2/spu/stock_update/"        // 多门店SPU库存同步(商品库存同步)
	SPU_SYNC_REQUEST_ADDRESS                = WEBSITE_URl + "poi/v2/spu/sync/"                // 多门店SPU同步(上线商品)

	//---------新商品同步-------- https://bytedance.feishu.cn/docx/doxcncpgP3vi7QtK4CPX7Ark1sd
	GOODS_TEMPLATE_GET_URL = WEBSITE_URl + "life/goods/template/get/"      // 查询商品模板
	GOODS_SAVE_URL         = WEBSITE_URl + "life/goods/product/save/"      // 创建/更新商品
	GOODS_DRAFT_GET_URL    = WEBSITE_URl + "life/goods/product/draft/get/" // 查询商品草稿数据
	GOODS_OPERATE_URL      = WEBSITE_URl + "life/goods/product/operate/"   // 上下架商品
)

// CurlGet 发送Get请求
func CurlGet(url, contentType, token string) (response string) {
	client := http.Client{Timeout: 5 * time.Second}
	resp, error := client.Get(url)
	resp.Header.Add("Content-Type", contentType)
	resp.Header.Add("access-token", token)
	if resp == nil {
		return ""
	}
	defer resp.Body.Close()
	if error != nil {
		panic(error)
	}

	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	response = result.String()
	return
}

// CurlPosts 发送POST请求(https)
func CurlPosts(url string, data, contentType, token string, types int) (content string) {
	var jsonStr []byte
	if types == 1 {
		jsonStr, _ = json.Marshal(data)
	} else if types == 2 {
		jsonStr = []byte(data)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("access-token", token)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	//跳过证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, error := client.Do(req)
	if resp == nil {
		return ""
	}
	if error != nil {
		panic(error)
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	content = string(result)
	return
}

// CurlPostsForm 另一种方式提交
func CurlPostsForm(Url string, data map[string]string, contentType, token string, types int) (content string) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	for key, val := range data {
		_ = writer.WriteField(key, val)
	}
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	req, err := http.NewRequest("POST", Url, payload)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	if err != nil {
		panic(err)
	}
	//跳过证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, error := client.Do(req)
	if resp == nil {
		return ""
	}
	if error != nil {
		panic(error)
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	content = string(result)
	return
}

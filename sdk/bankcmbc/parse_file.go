package bankcmbc

import (
	"errors"
	"root/src/helper/htime"
	"strconv"
	"strings"
)

// ParseFile 解析文件
func ParseFile(fileContent string, pubKey string) (rows []*ReadRow, err error) {
	begTime := htime.NowF14() // 导入批次开始时间
	//endTime := t14.Time(time.Now().Local().AddDate(0, 1, -1)).String()
	fileContent = strings.ReplaceAll(fileContent, "\r\n", "\n")
	fileContent = strings.TrimSpace(fileContent)

	// 根据=====拆分文件
	var ss = strings.Split(fileContent, "=====")
	var f1 = strings.TrimSpace(strings.Trim(ss[0], "\n"))
	var f2 = strings.TrimSpace(strings.Trim(ss[1], "\n"))
	err = RsaVery(f1, f2, pubKey)
	if err != nil {
		err = errors.New("文件签名验证失败 " + err.Error())
		return
	}

	// 变形客户号,加密后手机号,批次号,券类型编号,券数量,有效期,券产品编号
	// 换行符拆分
	var rr = strings.Split(f1, "\n")
	for i, r := range rr {
		if i == 0 {
			continue // 忽略标题行
		}
		var cols = strings.Split(r, ",")
		tel := strings.TrimSpace(cols[1]) // 手机号是sm3摘要

		openid := strings.TrimSpace(cols[0])
		if openid == "" {
			openid = tel
		}
		num, _ := strconv.ParseInt(strings.TrimSpace(cols[4]), 10, 64)

		actId := ""
		if len(cols) > 7 {
			actId = strings.TrimSpace(cols[7])
		}

		//变形客户号,加密后手机号,批次号,券类型编号,券数量
		rows = append(rows, &ReadRow{
			Openid:  strings.TrimSpace(openid),
			Tel:     strings.TrimSpace(tel),
			BatchId: strings.TrimSpace(cols[2]),
			ProdId:  strings.TrimSpace(cols[3]),
			Num:     num,
			BegTime: begTime,
			EndTime: strings.ReplaceAll(strings.TrimSpace(cols[5]), "-", "") + "235959",
			ActId:   actId,
		})
	}
	return
}

type ReadRow struct {
	Openid  string `json:"openid"`
	Tel     string `json:"tel"`
	BatchId string `json:"batchId"`
	ProdId  string `json:"prodId"`
	Num     int64  `json:"num"`
	BegTime string `json:"begTime"`
	EndTime string `json:"endTime"`
	ActId   string `json:"actId"`
}

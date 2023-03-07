package ldpush

import (
	"testing"
)

func TestPush(t *testing.T) {
	const SN = "21730N536628"
	const SN2 = "21730N536627"

	err := IPush(cfg(), &PushParams{
		MsgId:    "12312311137",
		Sn:       SN2, //WSD201231002
		Content:  "宁波放心充到账0.00元，放心消费有保障",
		TotalFee: "1000000.91",
	})
	if err != nil {
		t.Error(err)
	}

}

func TestPush2(t *testing.T) {
	/*err := ldpush.IPush(config.Now().Liandiyx, &ldpush.PushParams{
		MsgId:    "12312311134",
		Sn:       SN2,
		Content:  "宁波放心充到账0.00元，放心消费有保障",
		TotalFee: "103.8",
	})
	if err != nil {
		t.Error(err)
	}*/

}

func TestPush3(t *testing.T) {
	/*err := ldpush.IPush(config.Now().Liandiyx, &ldpush.PushParams{
		MsgId:    "123123111347",
		Sn:       SN2,
		Content:  "如果下载文件出现乱码，请选择另外一个 34.34 office 版本选项进行下载。",
		TotalFee: "103.8",
	})
	if err != nil {
		t.Error(err)
	}*/

}

// 生产验证
func TestPush4Prd(t *testing.T) {
	err := IPush(cfg1(), &PushParams{
		MsgId:    "123011231111211",
		Sn:       "LDHORN020101101100999", //LDHORN020101101100999         WSD201231002
		Content:  "宁波放心充到账0.00元，放心消费有保障",
		TotalFee: "12.9",
	})
	if err != nil {
		t.Error(err)
	}
}

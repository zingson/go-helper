package metro

import (
	"encoding/json"
	"testing"
)

// 根据手机号开票测试
func TestTicketOpen(t *testing.T) {
	tc, err := TicketOpen(getUpApp(), "13611703040", "GT021", Rand32())
	if err != nil {
		t.Error(err)
	}
	tbytes, _ := json.Marshal(tc)
	t.Log(string(tbytes))
	t.Log("success..............")
}

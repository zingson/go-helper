package bytedanceapp

import (
	"testing"
)

func TestGetPhoneNumber(t *testing.T) {

	r, err := GetPhoneNumber("tta549d3f43a70e97401", "Gaejl3YLTDMIFiKfwtfEtw==", "aa+8EF/YNczEJdpgwLJ/Ew==", "9WwZwbJrfkFoauRwDNEp7g7DMrAjcM5vk8BQHo8eVq/tJSL91eN4VcWTItflw7qx8JbihFZXg7rbJPwhMm+WH+Varx0k86MAyBN/5Z02CJKSS751X4o5jDXPR+IUHmbtAfB94t/x+yge16FjIG+RD6917xlaHPfxAvWDKdUCusgNwcOVniM9EjVb6cbZtsTCxDYsftib92Oa/ceu1hWiGQ==")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(r.JSON())

}

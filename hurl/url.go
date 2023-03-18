package hurl

import "strings"

// Splice 链接参数拼接
func Splice(url string, params map[string]string) string {
	if params == nil {
		params = make(map[string]string)
	}

	if strings.Contains(url, "#") {
		if strings.Contains(strings.Split(url, "#")[1], "?") {
			url = url + "&"
		} else {
			url = url + "?"
		}
	} else {
		if strings.Contains(url, "?") {
			url = url + "&"
		} else {
			url = url + "?"
		}
	}
	query := ""
	for k, v := range params {
		if v == "" {
			continue
		}
		query += "&" + k + "=" + v
	}
	if len(query) == 0 {
		return url
	}
	return url + query[1:]
}

package hurl

import "strings"

// Splice 链接参数拼接, 如果由#号拼接在#号后的链接上
func Splice(url string, params map[string]string) string {
	if params == nil {
		params = make(map[string]string)
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

	return url + query[1:]
}

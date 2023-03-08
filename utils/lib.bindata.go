package utils

import (
	"dmkt/src/tpl"
	"net/http"
)

// BinDataWriterHTML 通过go-bindata输出html模板
func BinDataWriterHTML(w http.ResponseWriter, name string, data interface{}) (err error) {
	bytes, err := tpl.Asset(name)
	if err != nil {
		return
	}
	err = WriterHTML(w, string(bytes), data)
	if err != nil {
		return
	}
	return
}

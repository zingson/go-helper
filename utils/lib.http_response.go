package utils

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"sync"
	"text/template"
)

var (
	_tpl sync.Map
)

// WriterHTML 渲染HTML模板
func WriterHTML(w http.ResponseWriter, htmlTemplate string, data interface{}) (err error) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var t *template.Template
	name := fmt.Sprintf("%x", md5.Sum([]byte(htmlTemplate)))
	tp, ok := _tpl.Load(name)
	if ok {
		t = tp.(*template.Template)
	} else {
		t, err = template.New(name).Parse(htmlTemplate)
		if err != nil {
			log.Error(err)
			return
		}
		_tpl.Store(name, t)
	}
	err = t.Execute(w, data)
	if err != nil {
		log.Error(err)
		return
	}
	return
}

// WriterJSON 输出JSON
func WriterJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	wbytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Write(wbytes)
}

// WriterRedirect 302 重定向
func WriterRedirect(w http.ResponseWriter, url string) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.Header().Set("Location", url)
	w.WriteHeader(302)
}

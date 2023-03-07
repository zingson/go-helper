package html_template

import (
	"bytes"
	"embed"
	"html/template"
)

const (
	DelimsLeft  = "{{"
	DelimsRight = "}}"
)

var (
	Tpl *template.Template
)

// ParseFS 解析模板
func ParseFS(funcMap template.FuncMap, fs embed.FS, patterns ...string) *template.Template {
	fm := template.FuncMap{
		"slot": Slot,
	}
	for k, v := range funcMap {
		fm[k] = v
	}
	t := template.Must(template.New("app").Delims(DelimsLeft, DelimsRight).Funcs(fm).ParseFS(fs, patterns...))
	Tpl = t
	return t
}

func Slot(name string, data any) template.HTML {
	var buf bytes.Buffer
	err := Tpl.ExecuteTemplate(&buf, name, data)
	if err != nil {
		panic(err)
	}
	return template.HTML(buf.String())
}

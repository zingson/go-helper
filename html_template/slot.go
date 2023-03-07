package html_template

import (
	"bytes"
	"embed"
	"html/template"
)

// Slot 模板实现插槽方式使用
func Slot1(fs embed.FS) func(name string, data any) template.HTML {
	return func(name string, data any) template.HTML {
		tBytes, err := fs.ReadFile(name)
		if err != nil {
			panic(err)
		}
		var buf bytes.Buffer
		err = template.Must(template.New(name).Parse(string(tBytes))).Execute(&buf, data)
		if err != nil {
			panic(err)
		}
		return template.HTML(buf.String())
	}
}

func Slot2(name string, data any) template.HTML {
	var buf bytes.Buffer
	err := Tpl.ExecuteTemplate(&buf, name, data)
	if err != nil {
		panic(err)
	}
	return template.HTML(buf.String())
}

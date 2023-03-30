package hstruct

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

/*
// Assign src的属性 赋值给 dst的属性 ，如果属性不存在则保持不变
func Assign(src interface{}, dst interface{}) {
	dstVal := reflect.ValueOf(dst).Elem()
	srcVal := reflect.ValueOf(src).Elem()
	srcType := srcVal.Type()
	for i := 0; i < srcVal.NumField(); i++ {
		name := srcType.Field(i).Name
		value := dstVal.FieldByName(name)
		if ok := value.IsValid(); ok {
			value.Set(reflect.ValueOf(srcVal.Field(i).Interface()))
		}
	}
}
*/

// Assign 使用gob实现结构体 src的属性 赋值给 dst的属性 ，如果属性不存在则保持不变
func Assign(src, dis any) {
	buff := new(bytes.Buffer)
	if err := gob.NewEncoder(buff).Encode(src); err != nil {
		panic(err)
	}
	if err := gob.NewDecoder(buff).Decode(dis); err != nil {
		panic(err)
	}
	return
}

type Struct struct {
}

func (o *Struct) Json() string {
	b, _ := json.Marshal(o)
	return string(b)
}

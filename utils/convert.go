package utils

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Int2String(data int) string {
	return strconv.Itoa(data)
}

func String2Int(data string) int {
	k, _ := strconv.Atoi(data)
	return k
}

func String2Int64(data string) int64 {
	k, _ := strconv.ParseInt(data, 10, 64)
	return k
}

func ConvertStructData(src interface{}, dst ...interface{}) error {
	jsonStr, err := json.Marshal(src)
	if err != nil {
		return err
	}

	for _, v := range dst {
		if err := json.Unmarshal(jsonStr, v); err != nil {
			return err
		}
	}

	return nil
}

func StructAssign(src interface{}, dst interface{}) {
	bVal := reflect.ValueOf(dst).Elem() //获取reflect.Type类型
	vVal := reflect.ValueOf(src).Elem() //获取reflect.Type类型
	vTypeOfT := vVal.Type()
	for i := 0; i < vVal.NumField(); i++ {
		// 在要修改的结构体中查询有数据结构体中相同属性的字段，有则修改其值
		name := vTypeOfT.Field(i).Name
		if ok := bVal.FieldByName(name).IsValid(); ok {
			bVal.FieldByName(name).Set(reflect.ValueOf(vVal.Field(i).Interface()))
		}
	}
}

func CopyOverride(dis, src interface{}) error {
	buff := new(bytes.Buffer)
	enc := gob.NewEncoder(buff)
	dec := gob.NewDecoder(buff)
	if err := enc.Encode(src); err != nil {
		return err
	}
	if err := dec.Decode(dis); err != nil {
		return err
	}
	return nil
}

func ConvertStruct2json(src interface{}) string {
	json, _ := json.Marshal(src)
	return string(json)
}

//StrContains 验证字符串是否包含字符串
//pattern 正则
func StrContains(pattern string, str string) bool {
	matched, _ := regexp.MatchString(pattern, str)
	return matched
}

func GetLastWeekSaturday() (weekMonday string) {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	thisWeekMonday := weekStartDate.Format("2006-01-02")
	TimeMonday, _ := time.Parse("2006-01-02", thisWeekMonday)
	lastWeekMonday := TimeMonday.AddDate(0, 0, -2)
	weekMonday = lastWeekMonday.Format("2006-01-02")
	return
}

func GetLastWeekMonday() (weekMonday string) {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	thisWeekMonday := weekStartDate.Format("2006-01-02")
	TimeMonday, _ := time.Parse("2006-01-02", thisWeekMonday)
	lastWeekMonday := TimeMonday.AddDate(0, 0, -7)
	weekMonday = lastWeekMonday.Format("2006-01-02")
	return
}

func GetWeek(week string) string {
	weekMap := map[string]string{
		"Monday":    "周一",
		"Tuesday":   "周二",
		"Wednesday": "周三",
		"Thursday":  "周四",
		"Friday":    "周五",
		"Saturday":  "周六",
		"Sunday":    "周天",
	}
	return weekMap[week]
}

// StringSplitAddIntNotRepeatElementIgnoreZero 在字符串分隔数组中
func StringSplitAddIntNotRepeatElementIgnoreZero(s, sep string, value int) string {
	// 在字符串最后拼上需要添加的数字
	s = fmt.Sprintf("%s%s%d", s, sep, value)
	// 转换为数组
	list := strings.Split(s, sep)
	// 创建一个临时map用来存储数组元素
	temp := make(map[int]bool)
	index := 0
	for _, v := range list {
		// 转换为数字
		number, err := strconv.Atoi(v)
		// 遍历数组元素，判断此元素是否已经存在map中
		if _, ok := temp[number]; err != nil || number == 0 || ok {
			list = append(list[:index], list[index+1:]...)
		}

		temp[number] = true
		index++
	}

	if len(list) == 0 {
		return ""
	}

	return strings.Join(list, sep)
}

// StringSliceRemoveRepeatElement 删除字符串切片中重复的和空的
func StringSliceRemoveRepeatElement(list ...string) []string {
	// 创建一个临时map用来存储数组元素
	temp := make(map[string]bool)
	index := 0
	for _, v := range list {
		// 遍历数组元素，判断此元素是否已经存在map中
		_, ok := temp[v]
		if ok || v == "" {
			list = append(list[:index], list[index+1:]...)
		} else {
			temp[v] = true
		}
		index++
	}
	return list
}

func RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}

//InArray 判断某一个值是否含在切片之中
func InArray(val interface{}, array interface{}) (exists bool) {
	exists = false

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				exists = true
				return
			}
		}
	}

	return
}

package goutils

import (
	"bytes"
	"encoding/json"
	"strings"
)

// Obj2Json 将go对象转换成json字符串
func Obj2Json(v interface{}) string {
	buf := bytes.NewBuffer([]byte{})
	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v)
	if err != nil {
		return ""
	}
	return strings.TrimRight(buf.String(), "\n")
}

// Obj2IndentJson 将go对象转换成json字符串(格式化)
func Obj2IndentJson(v interface{}) string {
	buf := bytes.NewBuffer([]byte{})
	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "\t")
	err := encoder.Encode(v)
	if err != nil {
		return ""
	}
	return strings.TrimRight(buf.String(), "\n")
}

// Json2Obj 将字符串反序列化到传入的对象指针中去
func Json2Obj(s string, v interface{}) error {
	data := []byte(s)
	return json.Unmarshal(data, v)
}

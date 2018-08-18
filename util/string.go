package util

import (
	"strconv"
	"unsafe"
)

//interface转为string
func interface2string(inter interface{}) string {
	tempStr := ""
	switch inter.(type) {
	case string:
		tempStr = inter.(string)
		break
	case float64:
		tempStr = strconv.FormatFloat(inter.(float64), 'f', -1, 64)
		break
	case int64:
		tempStr = strconv.FormatInt(inter.(int64), 10)
		break
	case int:
		tempStr = strconv.Itoa(inter.(int))
		break
	}
	return tempStr
}

//类型转换  string to bytes
func str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

//类型转换  bytes to string
func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

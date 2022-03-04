package utils

import (
	"unicode/utf8"

	"github.com/asaskevich/govalidator"
)

// Diff 基于切片排除切片中的数据。返回已排除数据的切片
func Diff(base, exclude []string) (result []string) {
	excludeMap := make(map[string]bool)
	for _, s := range exclude {
		excludeMap[s] = true
	}
	for _, s := range base {
		if !excludeMap[s] {
			result = append(result, s)
		}
	}
	return result
}

// Unique 去除切片重复数据
func Unique(ss []string) (result []string) {
	smap := make(map[string]bool)
	for _, s := range ss {
		smap[s] = true
	}
	for s := range smap {
		result = append(result, s)
	}
	return
}

// CamelCaseToUnderscore 驼峰转为下划线
func CamelCaseToUnderscore(str string) string {
	return govalidator.CamelCaseToUnderscore(str)
}

// UnderscoreToCamelCase 下换线转驼峰
func UnderscoreToCamelCase(str string) string {
	return govalidator.UnderscoreToCamelCase(str)
}

// FindString 找字符串在切片中的位置
func FindString(array []string, str string) int {
	for index, s := range array {
		if str == s {
			return index
		}
	}
	return -1
}

// StringIn 字符串是否存在切片中
func StringIn(str string, array []string) bool {
	return FindString(array, str) > -1
}

// Reverse 反转字符串
func Reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

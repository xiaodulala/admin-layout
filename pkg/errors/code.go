package errors

import (
	"fmt"
	"net/http"
	"sync"
)

var (
	unknownCoder defaultCoder = defaultCoder{
		C:    0,
		HTTP: http.StatusInternalServerError,
		Ext:  "An internal server error occurred",
		Ref:  "http://github.com/xiaodulala/component-tools/pkg/errors/README.md",
	}
)

// Coder 错误码信息的接口
type Coder interface {
	// HTTPStatus 业务码应该对应的HTTP状态码
	HTTPStatus() int
	// String 外部用户查看的错误文本信息
	String() string
	// Reference 详细的参考文档
	Reference() string
	// Code 错误码
	Code() int
}

type defaultCoder struct {
	C    int    // 错误码
	HTTP int    // http状态码
	Ext  string //暴露的错误信息
	Ref  string //引用文档
}

// HTTPStatus 返回http状态码
func (d defaultCoder) HTTPStatus() int {
	if d.HTTP == 0 {
		return 500
	}
	return d.HTTP
}

// String 返回错误信息
func (d defaultCoder) String() string {
	return d.Ext
}

// Reference 返回参考引用
func (d defaultCoder) Reference() string {
	return d.Ref
}

// Code 返回业务码
func (d defaultCoder) Code() int {
	return d.C
}

// codes 保存错误码的map
var codes = map[int]Coder{}
var codeMux = &sync.Mutex{}

// register 注册业务错误码。重复注册将覆盖
//nolint:deadcode,unused
func register(coder Coder) {
	if coder.Code() == 0 {
		panic("code `0` is reserved  as unknownCode error code")
	}

	codeMux.Lock()
	defer codeMux.Unlock()

	codes[coder.Code()] = coder
}

// mustRegister 注册业务错误码。重复注册将报错
func mustRegister(coder Coder) {
	if coder.Code() == 0 {
		panic("code `0` is reserved  as unknownCode error code")
	}

	codeMux.Lock()
	defer codeMux.Unlock()

	if _, ok := codes[coder.Code()]; ok {
		panic(fmt.Sprintf("code: %d already exist", coder.Code()))
	}

	codes[coder.Code()] = coder
}

// ParseCoder 解析错误为*withCode
func ParseCoder(err error) Coder {
	if err == nil {
		return nil
	}

	if v, ok := err.(*withCode); ok {
		if coder, ok := codes[v.code]; ok {
			return coder
		}
	}

	return unknownCoder
}

// IsCode 报告指定的错误链中是否包含指定的code码
func IsCode(err error, code int) bool {
	if v, ok := err.(*withCode); ok {
		if v.code == code {
			return true
		}

		if v.cause != nil {
			return IsCode(v.cause, code)
		}

		return false
	}

	return false
}

func init() {
	codes[unknownCoder.Code()] = unknownCoder
}

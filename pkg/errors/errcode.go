package errors

import "github.com/novalagung/gubrak"

// ErrCode 业务错误码
type ErrCode struct {
	C    int
	HTTP int
	Ext  string
	Ref  string
}

var _ Coder = &ErrCode{}

// HTTPStatus 错误码对应的http code
func (coder ErrCode) HTTPStatus() int {
	if coder.HTTP == 0 {
		return 500
	}
	return coder.HTTP
}

// String 对外错误信息
func (coder ErrCode) String() string {
	return coder.Ext
}

// Reference 参考文档
func (coder ErrCode) Reference() string {
	return coder.Ref
}

// Code 业务错误码
func (coder ErrCode) Code() int {
	return coder.C
}

// RegisterErrCoder 注册错误码
func RegisterErrCoder(code int, httpStatus int, message string, refs ...string) {
	found, _ := gubrak.Includes([]int{200, 400, 401, 403, 404, 500}, httpStatus)
	if !found {
		panic("http code not in `200, 400, 401, 403, 404, 500`")
	}

	var reference string
	if len(refs) > 0 {
		reference = refs[0]
	}

	coder := &ErrCode{
		C:    code,
		HTTP: httpStatus,
		Ext:  message,
		Ref:  reference,
	}

	mustRegister(coder)
}

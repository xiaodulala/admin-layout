package errors

import (
	"fmt"
	"io"
)

type withCode struct {
	err    error // error错误
	code   int   // 业务错误码
	cause  error // 根因错误
	*stack       //堆栈信息
}

// Error 实现error接口。自定义withCode错误类型
func (w *withCode) Error() string {
	return fmt.Sprintf("%v", w)
}

// Cause 返回包装错误的根因
func (w *withCode) Cause() error {
	return w.cause
}

// Unwrap 错误链解析
func (w *withCode) Unwrap() error {
	return w.cause
}

// WithCode 生成带业务错误码的错误根因
func WithCode(code int, format string, args ...interface{}) error {
	return &withCode{
		err:   fmt.Errorf(format, args...),
		code:  code,
		stack: callers(),
	}
}

// WrapC 生成带业务错误码,且包装根因的新错误
func WrapC(err error, code int, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	return &withCode{
		err:   fmt.Errorf(format, args...),
		code:  code,
		cause: err,
		stack: callers(),
	}
}

// Cause 返回错误链中的根因
func Cause(err error) error {
	type causer interface {
		Cause() error
	}

	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}
		if cause.Cause() == nil {
			break
		}

		err = cause.Cause()
	}

	return err
}

// New 生成根因的一种方式 无格式化
func New(message string) error {
	return &fundamental{
		msg:   message,
		stack: callers(),
	}
}

// Errorf 生成根因的一种方式 格式化错误
func Errorf(format string, args ...interface{}) error {
	return &fundamental{
		msg:   fmt.Sprintf(format, args...),
		stack: callers(),
	}
}

// fundamental 根本原因 包含错误信息和堆栈。
type fundamental struct {
	msg string
	*stack
}

func (f *fundamental) Error() string {
	return f.msg
}

func (f *fundamental) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			_, _ = io.WriteString(s, f.msg)

			f.stack.Format(s, verb)
			return
		}
		fallthrough
	case 's':
		_, _ = io.WriteString(s, f.msg)
	case 'q':
		_, _ = fmt.Fprintf(s, "%q", f.msg)
	}
}

// WithStack 被调用时在调用点记录有一个有堆栈跟踪的错误
func WithStack(err error) error {
	if err == nil {
		return nil
	}
	if e, ok := err.(*withCode); ok {
		return &withCode{
			err:   e.err,
			code:  e.code,
			cause: err,
			stack: callers(),
		}
	}

	return &withStack{
		error: err,
		stack: callers(),
	}
}

// withStack 包含堆栈的错误类型
type withStack struct {
	error
	*stack
}

func (w *withStack) Cause() error {
	return w.error
}

func (w *withStack) Unwrap() error {
	if e, ok := w.error.(interface{ Unwrap() error }); ok {
		return e.Unwrap()
	}
	return w.error
}

func (w *withStack) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v", w.Cause()) //nolint: errcheck
			w.stack.Format(s, verb)
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, w.Error()) //nolint: errcheck
	case 'q':
		fmt.Fprintf(s, "%q", w.Error()) //nolint: errcheck
	}
}

// Wrap 返回一个包含堆栈跟踪的错误，并且支持附加额外信息。
// 如果根错误为withCode类型。新的推展错误中也会包含。
func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}

	if e, ok := err.(*withCode); ok {
		return &withCode{
			err:   fmt.Errorf(message),
			code:  e.code,
			cause: err,
			stack: callers(),
		}
	}

	err = &withMessage{
		cause: err,
		msg:   message,
	}
	return withStack{
		error: err,
		stack: callers(),
	}
}

// Wrapf 支持格式化错误
func Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	if e, ok := err.(*withCode); ok {
		return &withCode{
			err:   fmt.Errorf(format, args...),
			code:  e.code,
			cause: err,
			stack: callers(),
		}
	}

	err = &withMessage{
		cause: err,
		msg:   fmt.Sprintf(format, args...),
	}
	return &withStack{
		err,
		callers(),
	}
}

// WithMessage 生成一个附件消息的错误
func WithMessage(err error, message string) error {
	if err == nil {
		return nil
	}
	return &withMessage{
		cause: err,
		msg:   message,
	}
}

// WithMessagef 生成一个附加消息错误，带格式化
func WithMessagef(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &withMessage{
		cause: err,
		msg:   fmt.Sprintf(format, args...),
	}
}

// withMessage 错误类型。附加消息
type withMessage struct {
	cause error
	msg   string
}

func (w *withMessage) Error() string {
	return w.msg
}

func (w *withMessage) Cause() error {
	return w.cause
}

func (w *withMessage) Unwrap() error {
	return w.cause
}
func (w *withMessage) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v\n", w.Cause())
			io.WriteString(s, w.msg) //nolint: errcheck
			return
		}
		fallthrough
	case 's', 'q':
		io.WriteString(s, w.Error()) //nolint: errcheck
	}
}

package zerrors

import "fmt"

type BusinessError struct {
	code Code
	err  error
}

func (e *BusinessError) ErrCode() Code {
	return e.code
}

func (e *BusinessError) ErrMsg() string {
	msg := e.code.String()
	if e.err != nil && e.err.Error() != "" {
		msg = fmt.Sprintf("%s[err:%s]", msg, e.err.Error())
	}
	return msg
}

func (e *BusinessError) Error() string {
	return e.ErrMsg()
}

func (e *BusinessError) Unwrap() error {
	return e.err
}

func NewBusinessError(code Code, err error) *BusinessError {
	return &BusinessError{
		code: code,
		err:  err,
	}
}

func NewBusinessErrorf(code Code, format string, a ...interface{}) *BusinessError {
	return &BusinessError{
		code: code,
		err:  fmt.Errorf(format, a),
	}
}

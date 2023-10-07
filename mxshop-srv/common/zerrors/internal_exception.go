package zerrors

import "fmt"

type InternalException struct {
	code Code
	err  error
}

func (e *InternalException) ErrCode() Code {
	return e.code
}

func (e *InternalException) ErrMsg() string {
	msg := e.code.String()
	if e.err != nil && e.err.Error() != "" {
		msg = fmt.Sprintf("%s[err:%s]", msg, e.err.Error())
	}
	return msg
}

func (e *InternalException) Error() string {
	return e.ErrMsg()
}

func (e *InternalException) Unwrap() error {
	return e.err
}

func NewInternalException(code Code, err error) *InternalException {
	return &InternalException{
		code: code,
		err:  err,
	}
}

func NewInternalExceptionf(code Code, format string, a ...interface{}) *InternalException {
	return &InternalException{
		code: code,
		err:  fmt.Errorf(format, a),
	}
}

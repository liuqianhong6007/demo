package internal

type WrapError struct {
	err error
	msg string
}

func (e WrapError) Error() string {
	msg := e.msg
	if e.err != nil {
		msg += "\n" + e.err.Error()
	}
	return msg
}

func (e WrapError) ExposeError() string {
	return e.msg
}

func wrapErrorFuncWithMsg(msg string) func(err error) WrapError {
	return func(err error) WrapError {
		return WrapError{err, msg}
	}
}

func wrapErrorFuncWithErr(err error) func(msg string) WrapError {
	return func(msg string) WrapError {
		return WrapError{err, msg}
	}
}

var (
	DatabaseError      = wrapErrorFuncWithMsg("internal database error")
	ParamParseError    = wrapErrorFuncWithMsg("internal param parse error")
	ParamValidateError = wrapErrorFuncWithErr(nil)
)

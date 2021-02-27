package internal

type WrapError struct {
	Err error
	Msg string
}

func (e *WrapError) Error() string {
	msg := e.Msg
	if e.Err != nil {
		msg += "\n" + e.Err.Error()
	}
	return msg
}

func (e *WrapError) ExposeError() string {
	return e.Msg
}

func wrapErrorFuncWithMsg(msg string) func(err error) *WrapError {
	return func(err error) *WrapError {
		if err == nil {
			return nil
		}
		return &WrapError{err, msg}
	}
}

func wrapErrorFuncWithErr(err error) func(msg string) *WrapError {
	return func(msg string) *WrapError {
		return &WrapError{err, msg}
	}
}

var (
	DatabaseError      = wrapErrorFuncWithMsg("internal database error")
	ParamParseError    = wrapErrorFuncWithMsg("internal param parse error")
	ParamValidateError = wrapErrorFuncWithErr(nil)
)

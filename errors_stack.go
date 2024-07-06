package errord

type errorStack struct {
	Err error
	Pc  uintptr
}

func newEmptyErrorStack() errorStack {
	return errorStack{}
}

func newErrorStack(err error, pc uintptr) errorStack {
	return errorStack{
		Err: err,
		Pc:  pc,
	}
}

func (e errorStack) Error() string {
	return e.Err.Error()
}

func (e errorStack) Unwrap() error {
	return e.Err
}

// implement ProgramCounterGetter interface
func (e errorStack) ProgramCounter() uintptr {
	return e.Pc
}

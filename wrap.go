package errord

import (
	"fmt"
	"runtime"
)

// Errorf is a wrapper for fmt.Errorf that also captures the stack trace.
func Errorf(format string, args ...any) error {
	return wrapStack(fmt.Errorf(format, args...))
}

func wrapStack(err error) errorStack {
	fpcs := make([]uintptr, 1)
	n := runtime.Callers(3, fpcs)
	if n == 0 {
		return newEmptyErrorStack()
	}
	return newErrorStack(err, fpcs[0])
}

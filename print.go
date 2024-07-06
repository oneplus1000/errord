package errord

import (
	"bytes"
	"fmt"
	"runtime"
	"slices"
)

func StackPrint(errVal any) {
	fmt.Print(StackString(errVal))
}

func StackPrintln(errVal any) {
	fmt.Println(StackString(errVal))
}

func Stack(errVal any) []CodeLine {
	if errVal == nil {
		return nil
	}
	var codelines []CodeLine
	for {
		uwpr, ok := errVal.(unwrpper)
		if !ok {
			break //cannot unwrap anymore
		}
		errVal = uwpr.Unwrap()
		if pcGetter, ok := uwpr.(ProgramCounterGetter); ok {
			pc := pcGetter.ProgramCounter()
			caller := runtime.FuncForPC(pc - 1)
			if caller == nil {
				continue
			}
			file, line := caller.FileLine(pc - 1)
			errMsgLocal := ""
			if val, ok := uwpr.(error); ok {
				errMsgLocal = val.Error()
			}
			codelines = append(codelines, CodeLine{
				file: file,
				line: line,
				msg:  errMsgLocal,
			})
		}

	}
	slices.Reverse(codelines)
	return codelines
}

func StackString(errVal any) string {
	if errVal == nil {
		return ""
	}

	codelines := Stack(errVal)
	var msg bytes.Buffer
	errMsgTop := fmt.Sprintf("%v", errVal)
	if errMsgTop != "" {
		msg.WriteString(errMsgTop)
		msg.WriteString("\n")
	}
	for i, cl := range codelines {
		if i > 0 {
			msg.WriteString("\n")
		}
		msg.WriteString(fmt.Sprintf("%s:%d", cl.file, cl.line))
	}
	return msg.String()
}

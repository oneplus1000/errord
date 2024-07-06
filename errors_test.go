package errord

import (
	"errors"
	"strings"
	"testing"
)

func TestError(t *testing.T) {
	err01 := errors.New("error 01")
	err02 := Errorf("error 02 : %w", err01)
	err03 := Errorf("error 03 : %w", err02)
	stk := StackString(err03)
	ok1 := strings.Contains(stk, "error 01")
	ok2 := strings.Contains(stk, "error 02")
	ok3 := strings.Contains(stk, "error 03")
	if !ok1 || !ok2 || !ok3 {
		t.Errorf("Error stack not found")
	}
	//fmt.Print("\n-----\n")
	//fmt.Print(stk)
	//fmt.Print("\n-----\n")
}

package errord

// ProgramCounterGetter is an interface that returns the program counter.
type ProgramCounterGetter interface {
	ProgramCounter() uintptr
}

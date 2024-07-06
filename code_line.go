package errord

// CodeLine is a struct that contains the file, line, and message of an error.
type CodeLine struct {
	file string
	line int
	msg  string
}

func (c CodeLine) File() string {
	return c.file
}

func (c CodeLine) Line() int {
	return c.line
}

func (c CodeLine) Msg() string {
	return c.msg
}

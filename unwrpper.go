package errord

type unwrpper interface {
	Unwrap() error
}

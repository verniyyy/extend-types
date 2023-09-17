package exterrors

import (
	"errors"
)

type Errors []error

func NewErrors() Errors {
	return make(Errors, 0)
}

func NewErrorsWithCap(size int) Errors {
	return make(Errors, 0, size)
}

func (errs *Errors) Add(e ...error) {
	if e == nil {
		return
	}
	*errs = append(*errs, e...)
}

// Join returns an error that wraps the self Errors.
// The error formats as the concatenation of the strings obtained
// by calling the Error method of each element of errs, with a newline
// between each string.
func (errs Errors) Join() error {
	if len(errs) == 0 {
		return nil
	}
	return errors.Join(errs...)
}

package list

import "fmt"

const (
	indexErrorMessageFormat = "index %d outside of list bounds: -%d...%d"
)

func IndexError(size uint, index int) error {
	return &indexError{
		size:  size,
		index: index,
	}
}

var _ error = indexError{}

type indexError struct {
	size  uint
	index int
}

func (e indexError) Error() string {
	return fmt.Sprintf(indexErrorMessageFormat, e.index, e.size, e.size-1)
}

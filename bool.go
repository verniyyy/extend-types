package types

import (
	"fmt"

	"github.com/verniyyy/extend-types/lib"
)

type Bool bool

func NewBool[T ~bool](b T) Bool {
	return Bool(b)
}

func (b Bool) String() string {
	return fmt.Sprintf("%t", bool(b))
}

func (b Bool) Ptr() *bool {
	return lib.Ptr(bool(b))
}

func (b Bool) And(v bool) bool {
	return bool(b) && v
}

func (b Bool) Or(v bool) bool {
	return bool(b) || v
}

func (b Bool) Not(v bool) bool {
	return !bool(b)
}

func (b Bool) XOr(v bool) bool {
	return (bool(b) || v) && !(bool(b) && v)
}

// Print with label.
// if label="example" then output:
// example: ${b}
func (b Bool) Print(label string) {
	fmt.Printf("%s: %v\n", label, b)
}

func (b Bool) CoreTypeName() string {
	return lib.TypeName(b)
}

func (b Bool) Validate() error {
	return nil
}

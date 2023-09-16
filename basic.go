package types

import (
	"fmt"

	"github.com/verniyyy/extend-types/lib"
)

type basic[T comparable] interface {
	fmt.Stringer
	Comparator[T]

	Read() T
	Validate() error
	Ptr() *T

	// Print with label.
	// if label="example" then output:
	// example: ${b.v}
	Print(label string)

	// CoreTypeName returns the type parameter name of a
	// basic type as a string.
	CoreTypeName() string
}

type b[T comparable] struct {
	v T
}

func newBasic[T comparable](v T) basic[T] {
	return &b[T]{
		v: v,
	}
}

func (b *b[T]) String() string {
	return fmt.Sprint(b.v)
}

func (b *b[T]) Read() T {
	return b.v
}

func (b *b[T]) Ptr() *T {
	return &b.v
}

func (b *b[T]) Equal(v T) bool {
	return b.v == v
}

func (b *b[T]) DeepEqual(v T) bool {
	return lib.DeepEqual(b.v, v)
}

// Print with label.
// if label="example" then output:
// example: ${b.v}
func (b *b[T]) Print(label string) {
	fmt.Printf("%s: %v\n", label, b.v)
}

func (b *b[T]) CoreTypeName() string {
	return lib.TypeName(b.v)
}

func (b *b[T]) Validate() error {
	return nil
}

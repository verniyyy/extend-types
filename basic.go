package types

import (
	"fmt"

	"github.com/verniyyy/extend-types/lib"
)

type basic[T comparable] interface {
	fmt.Stringer
	Read() T
	Ptr() *T
	Print()
	Println()
	PrimitiveTypeName() string
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

func (b *b[T]) Println() {
	fmt.Println(b.v)
}

func (b *b[T]) Print() {
	fmt.Print(b.v)
}

func (b *b[T]) PrimitiveTypeName() string {
	return lib.TypeName(b.v)
}

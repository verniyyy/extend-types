package types

import (
	"fmt"

	"github.com/verniyyy/extend-types-go/lib"
	"golang.org/x/exp/constraints"
)

type basic[T constraints.Ordered] struct {
	v T
}

func newBasic[T constraints.Ordered](v T) basic[T] {
	return basic[T]{
		v: v,
	}
}

func (b basic[T]) Value() T {
	return b.v
}

func (b basic[T]) String() string {
	return fmt.Sprint(b.v)
}

func (b basic[T]) Ptr() *T {
	return &b.v
}

func (b basic[T]) Equal(v T) bool {
	return b.v == v
}

func (b basic[T]) DeepEqual(v T) bool {
	return lib.DeepEqual(b.v, v)
}

// Print with label.
// if label="example" then output:
// example: ${b.v}
func (b basic[T]) Print(label string) {
	fmt.Printf("%s: %v\n", label, b.v)
}

func (b basic[T]) CoreTypeName() string {
	return lib.TypeName(b.v)
}

func (b basic[T]) Validate() error {
	return nil
}

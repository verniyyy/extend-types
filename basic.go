package types

import (
	"fmt"

	"github.com/YasutomoNakajima/extend-types/types/lib"
)

var ErrFreezed = fmt.Errorf("can't set value to freezed type")

type basic[T comparable] interface {
	read[T]
	write[T]
	constable
}

type read[T any] interface {
	fmt.Stringer
	Get() T
	Ptr() *T
	Print()
	PrintWithPrefix(string)
	PrimitiveTypeName() string
}

type write[T any] interface {
	Set(T) error
}

type constable interface {
	Freeze()
	IsFreezed() bool
}

type b[T comparable] struct {
	v         T
	isFreezed bool
}

func newBasic[T comparable](v T) basic[T] {
	return &b[T]{
		v:         v,
		isFreezed: false,
	}
}

func (b *b[T]) String() string {
	return fmt.Sprintf("%v", b.v)
}

func (b *b[T]) Get() T {
	return b.v
}

func (b *b[T]) Set(v T) error {
	if b.isFreezed {
		return ErrFreezed
	}

	b.v = v
	return nil
}

func (b *b[T]) Ptr() *T {
	return &b.v
}

func (b *b[T]) PrintWithPrefix(prefix string) {
	fmt.Print(prefix)
	b.print()
}

func (b *b[T]) Print() {
	b.print()
}

func (b *b[T]) print() {
	fmt.Println(b.v)
}

func (b *b[T]) PrimitiveTypeName() string {
	return lib.TypeName(b.v)
}

func (b *b[T]) Freeze() {
	b.isFreezed = true
}

func (b *b[T]) IsFreezed() bool {
	return b.isFreezed
}

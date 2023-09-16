package types

import (
	"golang.org/x/exp/constraints"
)

type Float[T constraints.Float] struct {
	basic[T]
}

func NewFloat[T constraints.Float](v T) Float[T] {
	return Float[T]{
		newBasic(v),
	}
}

func (f *Float[T]) Float32() float32 {
	return float32(f.Read())
}

func (f *Float[T]) Float64() float64 {
	return float64(f.Read())
}

func (f *Float[T]) Str() *Str {
	return fromStringer(f)
}

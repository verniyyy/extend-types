package types

import (
	"math"

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
	return float32(f.value())
}

func (f *Float[T]) Float64() float64 {
	return float64(f.value())
}

func (f *Float[T]) Ceil() T {
	return T(math.Ceil(f.Float64()))
}

func (f *Float[T]) Floor() T {
	return T(math.Floor(f.Float64()))
}

func (f *Float[T]) Round() T {
	return T(math.Round(f.Float64()))
}

func (f *Float[T]) Abs() T {
	return T(math.Abs(f.Float64()))
}

func (f *Float[T]) IsNegative() bool {
	return f.value() < 0
}

func (f *Float[T]) IsPositive() bool {
	return f.value() > 0
}

func (f *Float[T]) IsZero() bool {
	return f.value() == 0.0
}

func (f *Float[T]) Str() *Str {
	return fromStringer(f)
}

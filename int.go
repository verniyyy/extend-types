package types

import "golang.org/x/exp/constraints"

type Int[T constraints.Integer] struct {
	basic[T]
}

func NewInt[T constraints.Integer](i T) *Int[T] {
	return &Int[T]{
		newBasic(i),
	}
}

func (i *Int[T]) Int() int {
	return int(i.Read())
}

func (i *Int[T]) Int8() int8 {
	return int8(i.Read())
}

func (i *Int[T]) Int16() int16 {
	return int16(i.Read())
}

func (i *Int[T]) Int32() int32 {
	return int32(i.Read())
}

func (i *Int[T]) Int64() int64 {
	return int64(i.Read())
}

func (i *Int[T]) Str() *Str {
	return fromStringer(i)
}

func (i *Int[T]) Float() *Float[float64] {
	return NewFloat(float64(i.Read()))
}

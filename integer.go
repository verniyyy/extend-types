package types

import "golang.org/x/exp/constraints"

type Integer[T constraints.Integer] struct {
	basic[T]
}

func NewInteger[T constraints.Integer](i T) Integer[T] {
	return Integer[T]{
		newBasic(i),
	}
}

func (i *Integer[T]) Int() int {
	return int(i.Read())
}

func (i *Integer[T]) Int8() int8 {
	return int8(i.Read())
}

func (i *Integer[T]) Int16() int16 {
	return int16(i.Read())
}

func (i *Integer[T]) Int32() int32 {
	return int32(i.Read())
}

func (i *Integer[T]) Int64() int64 {
	return int64(i.Read())
}

func (i *Integer[T]) Str() *Str {
	return fromStringer(i)
}

func (i *Integer[T]) Float() Float[float64] {
	return NewFloat(float64(i.Read()))
}

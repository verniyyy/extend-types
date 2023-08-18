package types

import (
	"fmt"

	"github.com/samber/lo"
	"github.com/verniyyy/extend-types/lib"
)

type List[T any] []T

func NewList[T any](size int) List[T] {
	return make(List[T], size)
}

func NewListFromSlice[T any](s []T) List[T] {
	l := make(List[T], len(s))
	copy(l, s)
	return l
}

func NewListWithDefault[T any](size int, defaultVal T) List[T] {
	l := make(List[T], size)
	for i := range l {
		l[i] = defaultVal
	}
	return l
}

func (l List[T]) Size() int {
	return len(l)
}

func (l *List[T]) Append(v T) {
	*l = append(*l, v)
}

func (l List[T]) At(i int) T {
	index := l.index(i)
	return l[index]
}

func (l *List[T]) Insert(i int, v T) {
	index := l.index(i)
	*l = append((*l)[:index], append([]T{v}, (*l)[index:]...)...)
}

func (l *List[T]) Remove(i int) {
	index := l.index(i)
	*l = append((*l)[:index], (*l)[index+1:]...)
}

func (l List[T]) index(i int) int {
	if i < 0 {
		return len(l) + i
	}
	return i
}

func (l *List[T]) Concat(list List[T]) {
	*l = append((*l), list...)
}

func (l List[T]) IsEmpty() bool {
	return len(l) == 0
}

func (l List[T]) Find(predicate func(item T) bool) (T, bool) {
	return lo.Find(l, predicate)
}

func (l List[T]) Filter(predicate func(item T, index int) bool) List[T] {
	return lo.Filter(l, predicate)
}

func (l List[T]) Each(predicate func(item *T)) List[T] {
	copy := make([]*T, len(l))
	for i, v := range l {
		copy[i] = &v
	}
	for _, item := range copy {
		predicate(item)
	}
	result := lo.Map(copy, func(item *T, _ int) T {
		return *item
	})
	return result
}

func (l List[T]) EachWithIndex(predicate func(item *T, index int)) List[T] {
	copy := make([]*T, len(l))
	for i, v := range l {
		copy[i] = &v
	}
	for i, item := range copy {
		predicate(item, i)
	}
	result := lo.Map(copy, func(item *T, _ int) T {
		return *item
	})
	return result
}

func (l List[T]) Reject(predicate func(item T, index int) bool) List[T] {
	return lo.Reject(l, predicate)
}

func (l List[T]) CountBy(predicate func(item T) bool) int {
	return lo.CountBy(l, predicate)
}

func (l List[T]) Print() {
	l.print()
}

func (l List[T]) PrintWithPrefix(prefix string) {
	fmt.Print(prefix)
	l.print()
}

func (l List[T]) print() {
	fmt.Printf("%s{\n", lib.TypeName(l))
	for _, v := range l {
		fmt.Printf("\t%+v,\n", v)
	}
	fmt.Println("}")
}

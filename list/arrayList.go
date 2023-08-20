package list

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/verniyyy/extend-types/lib"
)

func NewArrayList[T any](size int) List[T] {

	return lib.Ptr(make(arrayList[T], size))
}

func NewArrayListFromSlice[T any](s []T) List[T] {
	l := make(arrayList[T], len(s))
	copy(l, s)
	return &l
}

func NewArrayListWithDefault[T any](size int, defaultVal T) List[T] {
	l := make(arrayList[T], size)
	for i := range l {
		l[i] = defaultVal
	}
	return &l
}

type arrayList[T any] []T

func (l arrayList[T]) Size() int {
	return len(l)
}

func (l *arrayList[T]) Add(v T) {
	*l = append(*l, v)
}

func (l arrayList[T]) At(i int) T {
	index := l.index(i)
	return l[index]
}

func (l *arrayList[T]) Insert(i int, v T) {
	index := l.index(i)
	*l = append((*l)[:index], append([]T{v}, (*l)[index:]...)...)
}

func (l *arrayList[T]) Remove(i int) {
	index := l.index(i)
	*l = append((*l)[:index], (*l)[index+1:]...)
}

func (l *arrayList[T]) Overwrite(i int, v T) {
	index := l.index(i)
	(*l)[index] = v
}

func (l *arrayList[T]) Concat(list List[T]) {
	*l = append((*l), *(list.(*arrayList[T]))...)
}

func (l arrayList[T]) IsEmpty() bool {
	return len(l) == 0
}

func (l arrayList[T]) Duplicate() List[T] {
	c := make(arrayList[T], len(l))
	copy(c, l)
	return &c
}

func (l arrayList[T]) DeepDuplicate() (List[T], error) {
	b := new(bytes.Buffer)
	enc := gob.NewEncoder(b)
	err := enc.Encode(l)
	if err != nil {
		return nil, err
	}

	d := make(arrayList[T], len(l))
	dec := gob.NewDecoder(b)
	err = dec.Decode(&d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

func (l arrayList[T]) Print() {
	l.print()
}

func (l arrayList[T]) PrintWithPrefix(prefix string) {
	fmt.Print(prefix)
	l.print()
}

func (l arrayList[T]) print() {
	fmt.Printf("%s{\n", lib.TypeName(l))
	for _, v := range l {
		fmt.Printf("\t%+v,\n", v)
	}
	fmt.Println("}")
}

func (l arrayList[T]) bounds() (from, to int) {
	return -1 * len(l), len(l) - 1
}

// index ...
func (l arrayList[T]) index(i int) int {
	from, to := l.bounds()
	if i < from || i >= to {
		panic(IndexError(uint(len(l)), i))
	}
	if i < 0 {
		return len(l) + i
	}
	return i
}

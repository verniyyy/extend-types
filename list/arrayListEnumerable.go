package list

import (
	"github.com/samber/lo"
	"github.com/verniyyy/extend-types-go/lib"
)

func (l arrayList[T]) Find(predicate func(item T) bool) (T, bool) {
	return lo.Find(l, predicate)
}

func (l arrayList[T]) Filter(predicate func(item T, index int) bool) List[T] {
	return (*arrayList[T])(lib.Ptr(lo.Filter(l, predicate)))
}

func (l arrayList[T]) Each(predicate func(item T)) {
	for _, item := range l {
		predicate(item)
	}
}

func (l arrayList[T]) EachWithIndex(predicate func(item T, index int)) {
	for i, item := range l {
		predicate(item, i)
	}
}

func (l arrayList[T]) Reject(predicate func(item T, index int) bool) List[T] {
	return (*arrayList[T])(lib.Ptr(lo.Reject(l, predicate)))
}

func (l arrayList[T]) Count(predicate func(item T) bool) int {
	return lo.CountBy(l, predicate)
}

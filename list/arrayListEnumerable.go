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

func (l arrayList[T]) Each(predicate func(item *T)) List[T] {
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
	return (*arrayList[T])(&result)
}

func (l arrayList[T]) EachWithIndex(predicate func(item *T, index int)) List[T] {
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
	return (*arrayList[T])(&result)
}

func (l arrayList[T]) Reject(predicate func(item T, index int) bool) List[T] {
	return (*arrayList[T])(lib.Ptr(lo.Reject(l, predicate)))
}

func (l arrayList[T]) Count(predicate func(item T) bool) int {
	return lo.CountBy(l, predicate)
}

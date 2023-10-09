package types

type Enumerator[T any, I comparable, R any] interface {
	Find(func(T) bool) (T, bool)
	Filter(func(T, I) bool) R
	Each(func(T))
	EachWithIndex(func(T, I))
	Reject(func(T, I) bool) R
	Count(func(T) bool) int
}

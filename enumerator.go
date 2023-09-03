package types

type Enumerator[T any, I comparable, R any] interface {
	Find(func(T) bool) (T, bool)
	Filter(func(T, I) bool) R
	Each(func(*T)) R
	EachWithIndex(func(*T, I)) R
	Reject(func(T, I) bool) R
	Count(func(T) bool) int
}

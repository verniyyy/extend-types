package types

type constants[T constable] struct {
	consts KeyValue[T]
}

type KeyValue[T constable] map[string]T

func NewConstants[T constable](m KeyValue[T]) *constants[T] {
	consts := make(KeyValue[T], len(m))
	for k, v := range m {
		consts[k] = v
		consts[k].Freeze()
	}
	return &constants[T]{
		consts: consts,
	}
}

func (c constants[T]) Value(key string) T {
	return c.consts[key]
}

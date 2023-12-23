package gople

// Gople ...
type Gople[T any] struct {
	V T
}

// Wrap ...
func Wrap[T any](v T) Gople[T] {
	return Gople[T]{
		V: v,
	}
}

// Unwrap ...
func (t Gople[T]) Unwrap() T {
	return t.V
}

// Gople2 ...
type Gople2[T1, T2 any] struct {
	V1 T1
	V2 T2
}

// Wrap2 ...
func Wrap2[T1, T2 any](v1 T1, v2 T2) Gople2[T1, T2] {
	return Gople2[T1, T2]{
		V1: v1,
		V2: v2,
	}
}

// Unwrap2 ...
func (t Gople2[T1, T2]) Unwrap2() (T1, T2) {
	return t.V1, t.V2
}

// Gople3 ...
type Gople3[T1, T2, T3 any] struct {
	V1 T1
	V2 T2
	V3 T3
}

// Wrap3 ...
func Wrap3[T1, T2, T3 any](v1 T1, v2 T2, v3 T3) Gople3[T1, T2, T3] {
	return Gople3[T1, T2, T3]{
		V1: v1,
		V2: v2,
		V3: v3,
	}
}

// Unwrap3 ...
func (t Gople3[T1, T2, T3]) Unwrap3() (T1, T2, T3) {
	return t.V1, t.V2, t.V3
}

// Gople4 ...
type Gople4[T1, T2, T3, T4 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
}

// Wrap4 ...
func Wrap4[T1, T2, T3, T4 any](v1 T1, v2 T2, v3 T3, v4 T4) Gople4[T1, T2, T3, T4] {
	return Gople4[T1, T2, T3, T4]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
	}
}

// Unwrap4 ...
func (t Gople4[T1, T2, T3, T4]) Unwrap4() (T1, T2, T3, T4) {
	return t.V1, t.V2, t.V3, t.V4
}

// Gople5 ...
type Gople5[T1, T2, T3, T4, T5 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
}

// Wrap5 ...
func Wrap5[T1, T2, T3, T4, T5 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5) Gople5[T1, T2, T3, T4, T5] {
	return Gople5[T1, T2, T3, T4, T5]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
	}
}

// Unwrap5 ...
func (t Gople5[T1, T2, T3, T4, T5]) Unwrap5() (T1, T2, T3, T4, T5) {
	return t.V1, t.V2, t.V3, t.V4, t.V5
}

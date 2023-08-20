package lib

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func DeepEqual[T any](x, y T) bool {
	nv := new(T)
	opt := cmp.AllowUnexported(*nv)
	return cmp.Equal(x, y, opt)
}

func DeepEqualIgnoreUnexported[T any](x, y T) bool {
	nv := new(T)
	opt := cmpopts.IgnoreUnexported(*nv)
	return cmp.Equal(x, y, opt)
}

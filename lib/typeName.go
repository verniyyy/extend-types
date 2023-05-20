package lib

import "reflect"

func TypeName[T any](v T) string {
	return reflect.TypeOf(v).Name()
}

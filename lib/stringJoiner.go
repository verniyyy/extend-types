package lib

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
)

func StringJoiner[T fmt.Stringer](sep string, s ...T) string {
	stringList := lo.Map(s, func(str T, _ int) string {
		return str.String()
	})
	return strings.Join(stringList, sep)
}

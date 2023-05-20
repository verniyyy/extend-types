package lib

import (
	"fmt"
	"testing"
)

func TestTypeName(t *testing.T) {
	i := 16
	n := TypeName(i)
	fmt.Printf("n: %v\n", n)
}

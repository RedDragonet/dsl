package dsl

import (
	"fmt"
	"testing"
)

func TestNewInterpreter(t *testing.T) {
	i := NewInterpreter("1+2")
	fmt.Println(i.expr())

	i = NewInterpreter("1+1")
	fmt.Println(i.expr())
}

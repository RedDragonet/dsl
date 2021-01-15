package dsl

import (
	"testing"
)

func Test_interpreter_expr(t *testing.T) {
	tests := []struct {
		name string
		text string
		want int
	}{
		{"1+1", "1+1", 2},
		{"1+2", "1+2", 3},

		{"11 +  11", "11 +  11", 22},
		{"11 +  22", "11 +  22", 33},

		{"11 -  11", "11 -  11", 0},
		{"11 -  22", "11 -  22", -11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := NewInterpreter(tt.text)
			got, _ := i.expr()
			if got != tt.want {
				t.Errorf("interpreter.expr() = %v, want %v", got, tt.want)
			}
		})
	}
}

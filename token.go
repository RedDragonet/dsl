package dsl

import "fmt"

const (
	INTEGER = iota + 1
	PLUS
	MINUS
	MULTIPLICATION
	DIVISION
	EOF
)

type token struct {
	iType int
	value string
}

func (t token) String() string {
	return fmt.Sprintf("Token({%d}, {%s})", t.iType, string(t.value))
}

func NewToken(iType int, value string) token {
	return token{
		iType: iType,
		value: value,
	}
}

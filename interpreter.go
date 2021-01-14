package dsl

import (
	"fmt"
	"strconv"
)

type interpreter struct {
	text         string
	pos          int
	currentToken token
}

func NewInterpreter(text string) interpreter {
	return interpreter{
		text: text,
	}
}

func (i *interpreter) getNextToken() (token, error) {
	text := i.text
	//到底
	if i.pos > len(text)-1 {
		return NewToken(EOF, ""), nil
	}

	currentChar := string(text[i.pos])
	if IsNumeric(currentChar) {
		token := NewToken(INTEGER, currentChar)
		i.pos++
		return token, nil
	}

	if currentChar == "+" {
		token := NewToken(PLUS, currentChar)
		i.pos++
		return token, nil
	}

	return token{}, fmt.Errorf("unknown token character %s", currentChar)
}
func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func (i *interpreter) eat(token int) error {
	var err error
	if i.currentToken.iType == token {
		i.currentToken, err = i.getNextToken()
	}

	return err
}

func (i *interpreter) expr() (int, error) {
	var err error
	i.currentToken, err = i.getNextToken()
	if err != nil {
		return 0, err
	}

	left := i.currentToken.value
	i.eat(INTEGER)

	_ = i.currentToken
	i.eat(PLUS)

	right := i.currentToken.value
	i.eat(INTEGER)

	leftNumber, err := strconv.Atoi(left)
	if err != nil {
		return 0, err
	}
	rightNumber, err := strconv.Atoi(right)
	if err != nil {
		return 0, err
	}
	return leftNumber + rightNumber, nil
}

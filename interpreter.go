package dsl

import (
	"fmt"
	"strconv"
)

//end of input
type interpreter struct {
	text         string
	pos          int
	currentChar  string
	currentToken token
}

func NewInterpreter(text string) interpreter {
	return interpreter{
		text:        text,
		currentChar: string(text[0]),
	}
}

func (i *interpreter) getNextToken() (token, error) {
	text := i.text
	//到底
	if i.pos > len(text)-1 {
		return NewToken(EOF, ""), nil
	}
	//本次循环主要为剔除空格
	for !i.isEoi() {
		if i.currentChar == " " {
			i.skipWhitespace()
			continue
		}

		if IsNumeric(i.currentChar) {
			token := NewToken(INTEGER, strconv.Itoa(i.integer()))
			return token, nil
		}

		if i.currentChar == "+" {
			i.advance()
			return NewToken(PLUS, "+"), nil
		}

		if i.currentChar == "-" {
			i.advance()
			return NewToken(MINUS, "-"), nil
		}

		if i.currentChar == "*" {
			i.advance()
			return NewToken(MULTIPLICATION, "*"), nil
		}

		if i.currentChar == "/" {
			i.advance()
			return NewToken(DIVISION, "/"), nil
		}
	}

	return token{}, fmt.Errorf("unknown token character %s", i.currentChar)
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

func (i *interpreter) skipWhitespace() {
	if !i.isEoi() && i.currentChar == " " {
		i.advance()
	}
}

// is end of input
func (i *interpreter) isEoi() bool {
	return i.currentChar == "EOI"
}

func (i *interpreter) integer() int {
	var result string
	for !i.isEoi() && IsNumeric(i.currentChar) {
		result += i.currentChar
		i.advance()
	}
	resultInt, _ := strconv.Atoi(result)
	return resultInt
}

func (i *interpreter) advance() {
	i.pos++
	if i.pos > len(i.text)-1 {
		i.currentChar = "EOI"
	} else {
		i.currentChar = string(i.text[i.pos])
	}
}

func (i *interpreter) expr() (int, error) {
	var err error
	var result int

	i.currentToken, err = i.getNextToken()
	if err != nil {
		return 0, err
	}

	for {
		if i.currentToken.iType == EOF {
			return result, nil
		}

		var left, right string

		if i.currentToken.iType == INTEGER {
			left = i.currentToken.value
			i.eat(INTEGER)
		}

		op := i.currentToken
		switch op.iType {
		case PLUS:
			i.eat(PLUS)
		case MINUS:
			i.eat(MINUS)
		case MULTIPLICATION:
			i.eat(MULTIPLICATION)
		case DIVISION:
			i.eat(DIVISION)
		}

		right = i.currentToken.value
		i.eat(INTEGER)

		var leftNumber int
		if left != "" {
			leftNumber, err = strconv.Atoi(left)
			if err != nil {
				return 0, err
			}
		} else {
			leftNumber = result
		}

		rightNumber, err := strconv.Atoi(right)
		if err != nil {
			return 0, err
		}

		switch op.iType {
		case PLUS:
			result = leftNumber + rightNumber
		case MINUS:
			result = leftNumber - rightNumber
		case MULTIPLICATION:
			result = leftNumber * rightNumber
		case DIVISION:
			result = leftNumber / rightNumber
		default:
			return 0, fmt.Errorf("unknow operator")
		}
	}
}

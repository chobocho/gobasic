package main

import (
	"fmt"
	"strconv"
)

func getType(c uint8) int {
	if c == '+' || c == '-' || c == '*' || c == '/' || c == '%' {
		return OPERAND
	}
	if c >= '0' && c <= '9' {
		return NUMBER
	}
	if c == 'P' {
		return PI
	}
	if c == ' ' || c == '\t' || c == '\n' || c == '\r' {
		return WHITESPACE
	}
	if c == '(' || c == ')' {
		return PARENTHESIS
	}
	return UNKNOWN
}

func getParenthesis(code string, currentIdx int) (int, Node) {
	node := Node{Value: string(code[currentIdx]), Type: LEFT_PARENTHESIS}
	if code[currentIdx] == ')' {
		node.Type = RIGHT_PARENTHESIS
	}
	currentIdx++
	return currentIdx, node
}

func getOperand(code string, currentIdx int) (int, Node) {
	var op int
	switch code[currentIdx] {
	case '+':
		op = PLUS
	case '-':
		op = MINUS
	case '*':
		op = MULTIPLY
	case '/':
		op = DIVIDE
	case '%':
		op = MOD
	}

	var node Node
	if op == DIVIDE && currentIdx+1 < len(code) && code[currentIdx+1] == '/' {
		node = Node{Value: "//", Type: QUOTIENT}
		currentIdx += 2
	} else {
		node = Node{Value: string(code[currentIdx]), Type: op}
		currentIdx++
	}
	return currentIdx, node
}

func getPI(code string, currentIdx int) (int, Node) {
	if currentIdx+2 >= len(code) || code[currentIdx+1] != 'I' {
		return currentIdx, Node{Value: "", Type: UNKNOWN}
	}

	node := Node{Value: "PI", Type: PI}
	currentIdx += 2
	return currentIdx, node
}

func getNumber(code string, currentIdx int) (int, Node) {
	nu := ""
	for currentIdx < len(code) && code[currentIdx] >= '0' && code[currentIdx] <= '9' {
		nu += string(code[currentIdx])
		currentIdx++
	}
	num, err := strconv.ParseInt(nu, 10, 64)
	if err != nil {
		panic(err)
	}
	node := Node{Value: num, Type: NUMBER}
	return currentIdx, node
}

func Scan(code string) *Tokenizer {
	tokenizer := Tokenizer{tokens: []Node{}, pos: 0}

	currentIdx := 0
	for currentIdx < len(code) {
		switch getType(code[currentIdx]) {
		case NUMBER:
			var node Node
			currentIdx, node = getNumber(code, currentIdx)
			tokenizer.tokens = append(tokenizer.tokens, node)
		case OPERAND:
			var node Node
			currentIdx, node = getOperand(code, currentIdx)
			tokenizer.tokens = append(tokenizer.tokens, node)
		case PARENTHESIS:
			var node Node
			currentIdx, node = getParenthesis(code, currentIdx)
			tokenizer.tokens = append(tokenizer.tokens, node)
		case PI:
			var node Node
			currentIdx, node = getPI(code, currentIdx)
			tokenizer.tokens = append(tokenizer.tokens, node)
		case WHITESPACE:
			currentIdx++
		case UNKNOWN:
			fmt.Printf("Unknown token at %d -> %c\n", currentIdx, code[currentIdx])
			return nil
		default:
			fmt.Printf("Never come to here!token at %d -> %c\n", currentIdx, code[currentIdx])
			return nil
		}
	}
	tokenizer.tokens = append(tokenizer.tokens, Node{Value: "END", Type: PROGRAM_END})
	return &tokenizer
}

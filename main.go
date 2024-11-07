package main

import (
	"container/list"
	"fmt"
	strconv "strconv"
)

type Node struct {
	Value interface{}
	Type  int
}

const (
	OPERAND = iota
	PLUS
	MINUS
	MULTIPLY
	DIVIDE
	MOD
	NUMBER
	LEFT_PARENTHESIS
	RIGHT_PARENTHESIS
	PARENTHESIS
	WHITESPACE
	PROGRAM_END
	UNKNOWN
)

func getType(u uint8) int {
	if u == '+' || u == '-' || u == '*' || u == '/' || u == '%' {
		return OPERAND
	}
	if u >= '0' && u <= '9' {
		return NUMBER
	}
	if u == ' ' || u == '\t' || u == '\n' || u == '\r' {
		return WHITESPACE
	}
	if u == '(' || u == ')' {
		return PARENTHESIS
	}
	return UNKNOWN
}

func getParenthesis(code string, currentIdx int) (int, Node) {
	node := Node{Value: string(code[currentIdx]), Type: PARENTHESIS}
	currentIdx++
	return currentIdx, node
}

func getOperand(code string, currentIdx int) (int, Node) {
	node := Node{Value: string(code[currentIdx]), Type: OPERAND}
	currentIdx++
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

func scan(code string) *list.List {
	tokens := list.New()
	currentIdx := 0
	for currentIdx < len(code) {
		switch getType(code[currentIdx]) {
		case NUMBER:
			var node Node
			currentIdx, node = getNumber(code, currentIdx)
			tokens.PushBack(node)
			break
		case OPERAND:
			var node Node
			currentIdx, node = getOperand(code, currentIdx)
			tokens.PushBack(node)
			break
		case PARENTHESIS:
			var node Node
			currentIdx, node = getParenthesis(code, currentIdx)
			tokens.PushBack(node)
			break
		case WHITESPACE:
			currentIdx++
			break
		case UNKNOWN:
			fmt.Printf("Unknown token: at %d %c\n", currentIdx, string(code[currentIdx]))
			return nil
		}
	}
	tokens.PushBack(Node{Value: "END", Type: PROGRAM_END})
	return tokens
}

func main() {
	fmt.Println("Go Basic V0.1\n")

	code := "(10+20)*30%7"
	tokens := scan(code)

	printTokens(tokens)

	fmt.Println("=== Finished ===")
}

func printTokens(tokens *list.List) {
	if tokens == nil {
		fmt.Println("No tokens")
		return
	}

	for e := tokens.Front(); e != nil; e = e.Next() {
		node := e.Value.(Node)
		switch node.Type {
		case NUMBER:
			fmt.Println(node.Value.(int64))
			break
		case OPERAND:
			fmt.Println(node.Value.(string))
			break
		case PARENTHESIS:
			fmt.Println(node.Value.(string))
			break
		}
	}
}

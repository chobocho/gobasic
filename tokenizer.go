package main

import (
	"fmt"
	"strings"
)

type Tokenizer struct {
	tokens []Node
	pos    int
}

func (tokens *Tokenizer) ToString() string {
	if tokens == nil || len(tokens.tokens) == 0 {
		return ""
	}

	var sb strings.Builder
	for _, node := range tokens.tokens {
		switch node.Type {
		case NUMBER:
			sb.WriteString(fmt.Sprintf("%v ", node.Value.(int64)))
		case PLUS, MINUS, MULTIPLY, DIVIDE, QUOTIENT, MOD:
			sb.WriteString(fmt.Sprintf("%v ", node.Value.(string)))
		case LEFT_PARENTHESIS, RIGHT_PARENTHESIS:
			sb.WriteString(fmt.Sprintf("%v ", node.Value.(string)))
		}
	}
	return sb.String()
}

func (tokens *Tokenizer) printTokens() {
	if tokens == nil || len(tokens.tokens) == 0 {
		fmt.Println("No tokens")
		return
	}

	for _, node := range tokens.tokens {
		switch node.Type {
		case NUMBER:
			fmt.Println(node.Value.(int64))
		case PI:
			fmt.Println("PI")
		case PLUS, MINUS, MULTIPLY, DIVIDE, QUOTIENT, MOD:
			fmt.Println(node.Value.(string))
		case LEFT_PARENTHESIS, RIGHT_PARENTHESIS:
			fmt.Println(node.Value.(string))
		case PROGRAM_END:
			fmt.Println("END")
		default:
			fmt.Printf("UnKnow Token: %v\n", node.Value)
		}
	}
}

func (tokens *Tokenizer) ParseExpression() Expression {
	if tokens == nil || len(tokens.tokens) == 0 {
		return nil
	}

	left := tokens.ParseTerm()

	if left == nil || tokens.pos >= len(tokens.tokens) {
		return nil
	}

	for tokens.tokens[tokens.pos].Type == PLUS ||
		tokens.tokens[tokens.pos].Type == MINUS {
		op := tokens.tokens[tokens.pos].Type
		tokens.pos++
		right := tokens.ParseTerm()
		left = &BinaryOperation{op: op, left: left, right: right}
	}

	return left
}

func (tokens *Tokenizer) ParseTerm() Expression {
	if tokens == nil || len(tokens.tokens) == 0 {
		return nil
	}
	left := tokens.ParseFactor()

	if left == nil || tokens.pos >= len(tokens.tokens) {
		return nil
	}

	for tokens.tokens[tokens.pos].Type == MULTIPLY ||
		tokens.tokens[tokens.pos].Type == DIVIDE ||
		tokens.tokens[tokens.pos].Type == QUOTIENT ||
		tokens.tokens[tokens.pos].Type == MOD {
		op := tokens.tokens[tokens.pos].Type
		tokens.pos++
		right := tokens.ParseFactor()
		left = &BinaryOperation{op: op, left: left, right: right}
	}

	return left
}

func (tokens *Tokenizer) ParseFactor() Expression {
	if tokens == nil || len(tokens.tokens) == 0 {
		return nil
	}

	if tokens.tokens[tokens.pos].Type == MINUS {
		tokens.pos++
		return &UnaryOperation{op: MINUS, value: tokens.ParsePrimary()}
	}

	return tokens.ParsePrimary()
}

func (tokens *Tokenizer) ParsePrimary() Expression {
	if tokens == nil || len(tokens.tokens) == 0 {
		return nil
	}

	token := tokens.tokens[tokens.pos]

	switch token.Type {
	case PI:
		tokens.pos++
		return &Number{floatValue: 3.1415926, Type: FLOAT}
	case NUMBER:
		tokens.pos++
		return &Number{value: token.Value.(int64)}
	case LEFT_PARENTHESIS:
		tokens.pos++
		expression := tokens.ParseExpression()
		tokens.pos++
		return expression
	default:
		return nil
	}
}

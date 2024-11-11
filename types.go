package main

import "fmt"

const (
	OPERAND = iota
	PLUS
	MINUS
	MULTIPLY
	DIVIDE
	QUOTIENT
	MOD
	NUMBER
	LEFT_PARENTHESIS
	RIGHT_PARENTHESIS
	PARENTHESIS
	WHITESPACE
	PI
	FLOAT
	PROGRAM_END
	UNKNOWN
)

type Node struct {
	Value interface{}
	Type  int
}

type Expression interface {
	Eval() float64
}

type Number struct {
	value      int64
	floatValue float64
	Type       int
}

func (n *Number) Eval() float64 {
	if n.Type == FLOAT {
		return n.floatValue
	}
	return float64(n.value)
}

type UnaryOperation struct {
	op    int
	value Expression
}

func (u *UnaryOperation) Eval() float64 {
	if u.op == MINUS {
		if u.value == nil {
			return 0
		}
		return -u.value.Eval()
	}
	return u.value.Eval()
}

type BinaryOperation struct {
	op    int
	left  Expression
	right Expression
}

func (b *BinaryOperation) Eval() float64 {
	if b == nil || b.left == nil || b.right == nil {
		return 0
	}
	left := b.left.Eval()
	right := b.right.Eval()

	switch b.op {
	case PLUS:
		return left + right
	case MINUS:
		return left - right
	case MULTIPLY:
		return left * right
	case DIVIDE:
		return left / right
	case QUOTIENT:
		return float64(int(left) / int(right))
	case MOD:
		return float64(int(left) % int(right))
	default:
		fmt.Println("Can not Eval: Unknown operator!")
	}
	return 0
}

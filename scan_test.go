package main

import (
	"testing"
)

func TestGetType(t *testing.T) {
	tests := []struct {
		name string
		arg  uint8
		want int
	}{
		{"Plus Operand", '+', OPERAND},
		{"Minus Operand", '-', OPERAND},
		{"Multiplication Operand", '*', OPERAND},
		{"Division Operand", '/', OPERAND},
		{"Modulus Operand", '%', OPERAND},
		{"Number", '5', NUMBER},
		{"Whitespace", ' ', WHITESPACE},
		{"Newline", '\n', WHITESPACE},
		{"Tab", '\t', WHITESPACE},
		{"Carriage return", '\r', WHITESPACE},
		{"Open Parenthesis", '(', PARENTHESIS},
		{"Close Parenthesis", ')', PARENTHESIS},
		{"PI", ')', PI},
		{"Unknown", '#', UNKNOWN},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getType(tt.arg); got != tt.want {
				t.Errorf("getType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalc(t *testing.T) {
	tests := []struct {
		code string
		want float64
	}{
		{"PI * 0", 0},
		{"3+4", 7},
		{"3+4 * 0", 3},
		{"PI * 10", 31.415926},
	}

	for _, tt := range tests {
		t.Run(tt.code, func(t *testing.T) {
			tokens := Scan(tt.code)
			ast := tokens.ParseExpression()
			if got := ast.Eval(); got != tt.want {
				t.Errorf("getType() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
(* EBNF for a calculator that supports +, -, *, /, %, (), and unary minus. *)

expression   ::= term { ("+" | "-") term }
term         ::= factor { ("*" | "/" | "%") factor }
factor       ::= primary | "-" primary
primary      ::= number | "(" expression ")"
number       ::= digit { digit }
digit        ::= "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9"
*/

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
	value int64
}

func (n *Number) Eval() float64 {
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
			break
		case PLUS, MINUS, MULTIPLY, DIVIDE, QUOTIENT, MOD:
			sb.WriteString(fmt.Sprintf("%v ", node.Value.(string)))
			break
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
			break
		case PLUS, MINUS, MULTIPLY, DIVIDE, QUOTIENT, MOD:
			fmt.Println(node.Value.(string))
			break
		case LEFT_PARENTHESIS, RIGHT_PARENTHESIS:
			fmt.Println(node.Value.(string))
			break
		case PROGRAM_END:
			break
		default:
			fmt.Printf("UnKnow Token: %v\n", node.Value)
			break
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
	if tokens.tokens[tokens.pos].Type == LEFT_PARENTHESIS {
		tokens.pos++
		tmpToken := tokens.ParseExpression()
		tokens.pos++
		return tmpToken
	}
	tokens.pos++
	if token.Type == NUMBER {
		return &Number{value: token.Value.(int64)}
	}
	return nil
}

func getType(c uint8) int {
	if c == '+' || c == '-' || c == '*' || c == '/' || c == '%' {
		return OPERAND
	}
	if c >= '0' && c <= '9' {
		return NUMBER
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
		break
	case '-':
		op = MINUS
		break
	case '*':
		op = MULTIPLY
		break
	case '/':
		op = DIVIDE
		break
	case '%':
		op = MOD
		break
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
			break
		case OPERAND:
			var node Node
			currentIdx, node = getOperand(code, currentIdx)
			tokenizer.tokens = append(tokenizer.tokens, node)
			break
		case PARENTHESIS:
			var node Node
			currentIdx, node = getParenthesis(code, currentIdx)
			tokenizer.tokens = append(tokenizer.tokens, node)
			break
		case WHITESPACE:
			currentIdx++
			break
		case UNKNOWN:
			fmt.Printf("Unknown token at %d -> %c\n", currentIdx, code[currentIdx])
			return nil
		}
	}
	tokenizer.tokens = append(tokenizer.tokens, Node{Value: "END", Type: PROGRAM_END})
	return &tokenizer
}

func main() {
	fmt.Println("Go Basic V0.1\nPress q! or exit to exit")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">> ")
		reader.Discard(reader.Buffered())
		code, _ := reader.ReadString('\n')
		if len(code) > 0 {
			code = code[:len(code)-1]
		}
		if code == "exit" || code == "q!" {
			break
		}
		tokens := Scan(code)
		ast := tokens.ParseExpression()
		if ast != nil {
			fmt.Printf("%v = %v\n", tokens.ToString(), ast.Eval())
		}
	}
	fmt.Println("=== Finished ===")
}

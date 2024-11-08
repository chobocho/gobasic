package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Go Basic V0.1\nPress q! or exit to exit")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">> ")
		reader.Discard(reader.Buffered())
		code, _ := reader.ReadString('\n')
		code = strings.TrimSpace(code)
		if code == "q!" || code == "exit" || code == "quit" {
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

package main

import (
	"fmt"

	"github.com/is-hoku/postfix_notation/cmd"
)

func main() {
	input := "a+b*c*(d+e)"
	output := cmd.Infix2Postfix(input)
	fmt.Println(output)
}

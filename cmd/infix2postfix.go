package cmd

import (
	"log"
)

func Infix2Postfix(infix string) string {
	a := rune('a')
	z := rune('z')
	postfix := ""
	stack := newStack()

	for _, char := range infix {
		switch {
		// 演算数はそのまま出力
		case rune(char) >= a && rune(char) <= z:
			postfix += string(rune(char))
		// 演算子の場合、スタックの最後にそれより優先順位が高い演算子があるなら pop していって終ったら自分を push する
		case char == '+':
			for {
				l := stack.last()
				if l == '*' || l == '/' || l == '+' || l == '-' {
					p := stack.pop()
					if !(p == '(' || p == ')') { // 括弧は後置記法の出力として不要
						postfix += string(p)
					}
				} else {
					break
				}
			}
			stack.push(rune(char))
		case char == '-':
			for {
				l := stack.last()
				if l == '*' || l == '/' || l == '+' || l == '-' {
					p := stack.pop()
					if !(p == '(' || p == ')') {
						postfix += string(p)
					}
				} else {
					break
				}
			}
			stack.push(rune(char))
		case char == '*':
			for {
				l := stack.last()
				if l == '*' || l == '/' {
					p := stack.pop()
					if !(p == '(' || p == ')') {
						postfix += string(p)
					}
				} else {
					break
				}
			}
			stack.push(rune(char))
		case char == '/':
			for {
				l := stack.last()
				if l == '*' || l == '/' {
					p := stack.pop()
					if !(p == '(' || p == ')') {
						postfix += string(p)
					}
				} else {
					break
				}
			}
			stack.push(rune(char))
		case char == '(':
			stack.push(rune(char))
		// 右括弧の場合、左括弧がくるまで pop する
		case char == ')':
			for {
				p := stack.pop()
				if !(p == '(') {
					if !(p == '(' || p == ')') {
						postfix += string(p)
					}
				} else {
					break
				}
			}
			stack.push(rune(char))
		default:
			log.Fatal("invalid operand is in stack")
		}
	}
	for !stack.isEmpty() {
		p := stack.pop()
		if !(p == '(' || p == ')') {
			postfix += string(p)
		}
	}
	return postfix
}

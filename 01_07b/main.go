package main

import (
	"flag"
	"log"
)

func isClosed(r rune) bool {
	return r == ')' || r == ']' || r == '}'
}

func isOpen(r rune) bool {
	return r == '(' || r == '[' || r == '{'
}

func balanced(left, right rune) (b bool) {
	b = left == '(' && right == ')' ||
		left == '[' && right == ']' ||
		left == '{' && right == '}'
	return
}

// isBalanced returns whether the given expression
// has balanced brackets.
func isBalanced(expr string) bool {
	// an expression has balanced brackets if
	// …it has no brackets
	// …for each closed braket B _of type T_ the _last_ open bracket found is of the same type
	openBracketsStack := []rune{}
	for _, r := range expr {
		if isOpen(r) {
			openBracketsStack = append(openBracketsStack, r) //push
		}
		if isClosed(r) {
			if len(openBracketsStack) == 0 {
				return false
			}
			// pop and check
			lastIdx := len(openBracketsStack) - 1
			lastOpenBracket := openBracketsStack[lastIdx]
			if !balanced(lastOpenBracket, r) {
				return false
			}
			openBracketsStack = openBracketsStack[:lastIdx]
		}
	}
	return len(openBracketsStack) == 0
}

// printResult prints whether the expression is balanced.
func printResult(expr string, balanced bool) {
	if balanced {
		log.Printf("%s is balanced.\n", expr)
		return
	}
	log.Printf("%s is not balanced.\n", expr)
}

func main() {
	expr := flag.String("expr", "", "The expression to validate brackets on.")
	flag.Parse()
	printResult(*expr, isBalanced(*expr))
}

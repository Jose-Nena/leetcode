package main

import (
	"fmt"
	"strings"
)

func matchParentheses(leftParentheses string, rightParentheses string) bool {
	switch leftParentheses {
	case "(":
		if rightParentheses == ")" {
			return true
		}
	case "[":
		if rightParentheses == "]" {
			return true
		}
	case "{":
		if rightParentheses == "}" {
			return true
		}
	}
	return false
}

func leftParenthesesStack(parentheses string, leftStack []string) ([]string, bool) {
	if parentheses == "(" ||
		parentheses == "[" ||
		parentheses == "{" {
		leftStack = append(leftStack, parentheses)
	}

	if parentheses == ")" ||
		parentheses == "]" ||
		parentheses == "}" {
		if !matchParentheses(leftStack[len(leftStack)-1], parentheses) {
			return leftStack, false
		} else {
			leftStack = leftStack[:len(leftStack)-1]
		}
	}

	return leftStack, true
}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	allParentheses := strings.Split(s, "")
	leftStack := []string{}
	match := false

	for _, p := range allParentheses {
		leftStack, match = leftParenthesesStack(p, leftStack)
		if !match {
			return false
		}
	}
	return true
}

func main() {
	s1 := "()"
	fmt.Println(isValid(s1)) // true

	s2 := "()[]{}"
	fmt.Println(isValid(s2)) // true

	s3 := "(]"
	fmt.Println(isValid(s3)) // false

	s4 := "(])]"
	fmt.Println(isValid(s4)) // false

	s5 := "{[]}"
	fmt.Println(isValid(s5)) // true
}

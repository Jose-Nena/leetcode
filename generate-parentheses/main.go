package main

import (
	"fmt"
)

func backtrack(left int, right int, cur string, res *[]string) {
	if left < 0 || right < 0 {
		return
	}
	if right < left {
		return
	}
	if left == 0 && right == 0 {
		*res = append(*res, cur)
		return
	}

	//left
	//add
	cur = cur + "("
	backtrack(left-1, right, cur, res)
	//remove
	cur = cur[:len(cur)-1]

	//right
	//add
	cur = cur + ")"
	backtrack(left, right-1, cur, res)
	//remove
	cur = cur[:len(cur)-1]
}

func generateParentheses(n int) []string {
	res := []string{}
	if n == 0 {
		return res
	}
	backtrack(n, n, "", &res)
	return res
}

func main() {
	fmt.Println(generateParentheses(2))
	fmt.Println(generateParentheses(3))
	fmt.Println(generateParentheses(4))
}

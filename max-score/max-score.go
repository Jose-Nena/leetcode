package main

import (
	"fmt"
)

func maxScore(s string) int {
	res := 0
	if s == "" {
		return res
	}
	left := 0
	if s[0] == '0' {
		left = 1
	}
	right := 0
	for i := 1; i < len(s); i++ {
		if s[i] == '1' {
			right++
		}
	}
	res = left + right

	for i := 1; i < len(s)-1; i++ {
		if s[i] == '0' {
			left++
		} else {
			right--
		}
	}
	temp := left + right
	if temp > res {
		res = temp
	}

	return res
}

func main() {
	fmt.Println(maxScore("011101"))
}

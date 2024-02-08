package main

import "fmt"

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{0}
	}

	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += carry
		carry = digits[i] / 10
		digits[i] = digits[i] % 10
	}

	if carry > 0 {
		digits = append([]int{1}, digits...)
	}

	return digits
}

func main() {
	arr := []int{9, 9, 9}
	fmt.Println(plusOne(arr))
}

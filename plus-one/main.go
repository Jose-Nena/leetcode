package main

import "fmt"

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{0}
	}

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		if digits[i] == 9 {
			digits[i] = 0
			if i == 0 {
				digits = append([]int{1}, digits...)
			}
		}
	}
	return digits
}

func main() {
	digits := []int{7, 9, 9}
	fmt.Println(plusOne(digits))
}

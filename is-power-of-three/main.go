package main

import "fmt"

func isPowerofThree(n int) bool {
	return isPowerOfN(n, 3)
}

func isPowerOfN(num, n int) bool {
	if num < 1 {
		return false
	}
	for num%n == 0 {
		num /= n
	}
	return num == 1
}

func main() {
	fmt.Println(isPowerofThree(27))
	fmt.Println(isPowerofThree(0))
	fmt.Println(isPowerofThree(9))
	fmt.Println(isPowerofThree(45))
	fmt.Println(isPowerofThree(81))
	fmt.Println(isPowerofThree(243))
	fmt.Println(isPowerofThree(249))
	fmt.Println(isPowerofThree(390))
	fmt.Println(isPowerofThree(1))
}

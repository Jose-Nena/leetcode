package main

import "fmt"

func isPowerofThree(n int) bool {

	// Fiz o for abaixo só pra praticar uma alternativa.
	// Mas o for mais abaixo é mais simples e melhor.
	j := n
	for i := j; i > 3; i = n {
		if n%3 != 0 {
			return false
		}
		n /= 3
	}

	//	for n > 3 {
	//		if n%3 != 0 {
	//			return false
	//		}
	//		n /= 3
	//	}

	return n == 3

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
}

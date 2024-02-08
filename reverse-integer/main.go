package main

import "fmt"

func reverse(x int) []int {
	res := x
	//i := 0
	inv := []int{}

	for res >= 10 {
		inv = append(inv, res % 10)
		res = res/10
		//if res < 10 {
		//	inv = append(inv, res)
		//}
	}
	return inv
}

func main() {
	i := 1235
	fmt.Println(reverse(i)) // output = 5321

	//i = -123
	//fmt.Println(i) // output = -321

	i = 120
	fmt.Println(reverse(i)) // output = 21
}

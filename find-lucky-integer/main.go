package main

import "fmt"

func findLucky(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}

	dic := make(map[int]int)
	for _, v := range arr {
		dic[v]++
	}

	res := -1
	for k, v := range dic {
		if k == v && v > res {
			res = v
		}
	}

	return res
}

func main() {

	a := []int{1, 2, 2, 2, 3, 3, 3}
	fmt.Println(findLucky(a))

}

package main

import "fmt"

func uniqueOccurrences(arr []int) bool {
	if len(arr) == 0 {
		return false
	}

	numOccurs := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		numOccurs[arr[i]]++
	}

	repetedOccurs := make(map[int]bool)
	for _, v := range numOccurs {
		if _, ok := repetedOccurs[v]; ok {
			return false
		}
		repetedOccurs[v] = true
	}

	return true
}

func main() {

	arr1 := []int{1, 2, 2, 1, 1, 3}
	fmt.Println(uniqueOccurrences(arr1))

	arr2 := []int{1, 2, 5}
	fmt.Println(uniqueOccurrences(arr2))

	arr3 := []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}
	fmt.Println(uniqueOccurrences(arr3))
}

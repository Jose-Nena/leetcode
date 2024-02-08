package main

import "fmt"

func uniqueOccurrences(arr []int) bool {
	if arr == nil || len(arr) == 0 {
		return true
	}

	dic := make(map[int]int)
	for _, v := range arr {
		dic[v]++
	}

	countDic := make(map[int]int)
	for _, v := range dic {
		_, ok := countDic[v]
		if ok {
			return false
		} else {
			countDic[v]++
		}
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

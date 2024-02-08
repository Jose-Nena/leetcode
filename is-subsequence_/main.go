package main

import (
	"fmt"
)

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s) > len(t) {
		return false
	}

	first := 0
	second := 0
	for first < len(s) && second < len(t) {
		if s[first] != t[second] {
			second++
		} else {
			first++
			second++
		}
	}
	if first < len(s) && second == len(t) { // I really don't get this if. It seems to me that all we need is to test if first == len(s)
		return false
	}
	return first == len(s)
}

func main() {
	var s string = "abc"
	var t string = "axbxcx"
	fmt.Println(isSubsequence(s, t))

	s = "abc"
	t = "axcxbx"
	fmt.Println(isSubsequence(s, t))

	s = "abcde"
	t = "arxsdfdsbxcxdreeswe"
	fmt.Println(isSubsequence(s, t))

}

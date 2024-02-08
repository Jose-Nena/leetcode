package main

import (
	"fmt"
	"strings"
)

func checkPattern(a []string, b []string) bool {
	dic := make(map[string]string)
	p := 0
	for p < len(a) {
		val, ok := dic[a[p]]
		if !ok {
			dic[a[p]] = b[p]
		} else {
			if val != b[p] {
				return false
			}
		}
		p++
	}
	return true
}

func wordPattern(pattern string, str string) bool {
	strs := strings.Split(str, " ")
	patterns := strings.Split(pattern, "")

	if len(patterns) != len(strs) {
		return false
	}

	return checkPattern(patterns, strs) && checkPattern(strs, patterns)
}

func main() {
	var p1 string = "abba"
	var s1 string = "dog cat cat dog"
	fmt.Println(wordPattern(p1, s1)) // true

	var p2 string = "abba"
	var s2 string = "dog cat cat fish"
	fmt.Println(wordPattern(p2, s2)) // false

	var p3 string = "aaaa"
	var s3 string = "dog cat cat dog"
	fmt.Println(wordPattern(p3, s3)) // false

	var p4 string = "aacc"
	var s4 string = "love love peace peace"
	fmt.Println(wordPattern(p4, s4)) // true

	var p5 string = "abcd"
	var s5 string = "dog cat bird bird"
	fmt.Println(wordPattern(p5, s5)) // false

	var p6 string = "abbc"
	var s6 string = "dog cat cat bird"
	fmt.Println(wordPattern(p6, s6)) // true

	var p7 string = "aada"
	var s7 string = "dog dog dog dog"
	fmt.Println(wordPattern(p7, s7)) // false
}

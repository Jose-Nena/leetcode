package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, str string) bool {
	patternStringPairs := make(map[rune]string)
	stringPatternPairs := make(map[string]rune)
	patterns := strings.Split(pattern, "")
	strs := strings.Split(str, " ")

	if len(patterns) != len(strs) {
		return false
	}

	for i, r := range pattern {
		patternStringPairs[r] = strs[i]
		stringPatternPairs[strs[i]] = r
	}

	for p, s := range patternStringPairs {
		pat := stringPatternPairs[s]
		if p != pat {
			return false
		}
	}

	for s, p := range stringPatternPairs {
		strg := patternStringPairs[p]
		if s != strg {
			return false
		}
	}

	return true
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

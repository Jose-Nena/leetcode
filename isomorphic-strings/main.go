package main

import (
	"fmt"
	"strings"
)

func checkIsIsomorphic(a []string, b []string) bool {
	dic := make(map[string]string)
	for i, r := range a {
		v, ok := dic[r]
		if !ok {
			dic[r] = b[i]
		} else {
			if v != b[i] {
				return false
			}
		}
	}
	return true
}

func isIsomorphic(s string, t string) bool {
	ss := strings.Split(s, "")
	st := strings.Split(t, "")

	if len(ss) != len(st) {
		return false
	}

	return checkIsIsomorphic(ss, st) && checkIsIsomorphic(st, ss)
}

func main() {
	s1, t1 := "egg", "add"
	fmt.Println(isIsomorphic(s1, t1)) // true

	s2, t2 := "foo", "bar"
	fmt.Println(isIsomorphic(s2, t2)) // false

	s3, t3 := "paper", "title"
	fmt.Println(isIsomorphic(s3, t3)) // true

	s5, t5 := "saturday", "robsinoh"
	fmt.Println(isIsomorphic(s5, t5)) // true

	s6, t6 := "paper", "titli"
	fmt.Println(isIsomorphic(s6, t6)) // false
}

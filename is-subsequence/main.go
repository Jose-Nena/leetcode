package main

import (
	"fmt"
	"strings"
)

func isSubsequence(s string, t string) bool {
	s_arr := strings.Split(s, "")
	t_arr := strings.Split(t, "")
	m := make(map[string]bool)
	//res := []string{}
	j := 0
	for _, c := range s_arr {
		for i := j; i < len(t); i++ {
			if t_arr[i] == string(c) {
				j = i + 1
				m[t_arr[i]] = true
				break
			}
		}
	}
	return len(m) == len(s_arr)
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

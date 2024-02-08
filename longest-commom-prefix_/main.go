package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}
	res := strs[0]
	for i := 1; i < len(strs); i++ {
		resLength := len(res)
		tempLength := len(strs[i])
		length := min(resLength, tempLength)
		j := 0
		for j < length {
			if res[j] != strs[i][j] {
				break
			}
			j++
		}
		res = res[:j]
		if len(res) == 0 {
			return ""
		}
	}
	return res
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))

	strs = []string{"dog", "reacecar", "car"}
	fmt.Println(longestCommonPrefix(strs) + "passei dog")

	strs = []string{"predisposto", "preparado", "prerrogativa, prego"}
	fmt.Println(longestCommonPrefix(strs))

	strs = []string{"prdisposto", "preparado", "prerrogativa, prego"}
	fmt.Println(longestCommonPrefix(strs))

	strs = []string{"pdisposto", "preparado", "prerrogativa, prego"}
	fmt.Println(longestCommonPrefix(strs))

	strs = []string{"disposto", "preparado", "prerrogativa, prego"}
	fmt.Println(longestCommonPrefix(strs) + "passei disposto")
}

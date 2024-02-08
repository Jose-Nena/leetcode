package main

import (
	"errors"
	"fmt"
)

func findLongestPrefix(shorterString string, longerString string, longestCommonPrefix *string) error {
	s := ""
	commonPrefix := ""
	if shorterString[0] != longerString[0] {
		return errors.New("não há prefixo comum entre todas as strings")
	}
	for j, r := range shorterString {
		s = string(r)
		if shorterString[j] == longerString[j] {
			//*longestCommonPrefix = *longestCommonPrefix + s
			commonPrefix += s
		} else {
			*longestCommonPrefix = commonPrefix
			return nil
		}
	}
	*longestCommonPrefix = commonPrefix
	return nil
}

func longestCommonPrefix(strs []string) string {

	firstString := strs[0]
	longestCommonPrefix := ""

	var err error
	for i, _ := range strs {
		if i == len(strs)-1 {
			continue
		}
		if len(firstString) <= len(strs[i+1]) {
			err = findLongestPrefix(firstString, strs[i+1], &longestCommonPrefix)
		} else {
			err = findLongestPrefix(strs[i+1], firstString, &longestCommonPrefix)
		}

		if err != nil {
			return ""
		}
	}

	return longestCommonPrefix

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

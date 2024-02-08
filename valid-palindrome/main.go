package main

import "fmt"

// number and letter only
// A == a
func isPalindrome(s string) bool {
	if s == "" {
		return true
	}

	left, right := 0, len(s)-1

	for left < right {
		for left < right && !isLetterOrNumber(s[left]) {
			left++
		}

		for left < right && !isLetterOrNumber(s[right]) {
			right--
		}

		if left <= len(s)-1 && right >= 0 {
			if toUpper(s[left]) != toUpper(s[right]) {
				return false
			}
		}
		left++
		right--
	}
	return true
}

func isLetterOrNumber(char byte) bool {
	// A-Z (65 to 90)		a-z (97 to 122)	  numbers (48 to 57)
	return isUpper(char) || isLower(char) || isNumber(char)
}

func isUpper(char byte) bool {
	return 65 <= char && char <= 90
}

func isLower(char byte) bool {
	return 97 <= char && char <= 122
}

func isNumber(char byte) bool {
	return 48 <= char && char <= 57
}

func toUpper(char byte) byte {
	if isLower(char) {
		return byte(char - 32)
	}
	return char

}

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))

	s = "race a car"
	fmt.Println(isPalindrome(s))

	s = "Socorram-me, subi no onibus em Marrocos."
	fmt.Println(isPalindrome(s))

	s = "civic, radar, civic."
	fmt.Println(isPalindrome(s))

	s = "kinnikinnik..."
	fmt.Println(isPalindrome(s))

	s = "Sore was I ere I saw Eros."
	fmt.Println(isPalindrome(s))

	s = "Never a foot too far, even...."
	fmt.Println(isPalindrome(s))

	s = "Euston saw I was not Sue."
	fmt.Println(isPalindrome(s))

}

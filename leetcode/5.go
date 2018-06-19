package main

import "fmt"

func longestPalindrome(s string) string {

	length := len(s)
	if length == 0 {
		return ""
	}

	if length == 1 {
		return s
	}

	longest := 0
	left := 0
	right := 0

	start := 0
	for start < length-1 {
		// [start+1, lenght-1]

		findNew := false
		for i := length - 1; i > start; i-- {
			if s[start] == s[i] {
				low := start
				high := i
				isPalindrome := true
				for low < high {
					if s[low] != s[high] {
						isPalindrome = false
						break
					}
					low++
					high--
				}

				if isPalindrome {
					// 是回文数
					subLength := i - start + 1
					if subLength > longest {
						findNew = true
						longest = subLength
						left = start
						right = i
					}
					break
				}

			}
		}

		_ = findNew
		//
		start++
	}

	return s[left : right+1]
}

func main() {
	fmt.Println(longestPalindrome("abcddcbaeabcddcb"))
}

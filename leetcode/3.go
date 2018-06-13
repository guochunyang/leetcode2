package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}

	if length == 1 {
		return 1
	}

	maxLength := 0

	start := 0
	for i := 1; i < length; i++ {

		//fmt.Println("i:", i, s[i], s[start:i-1])
		// 在 [start, i-1] 中 倒序寻找 s[i]

		findNewSub := false
		for j := i - 1; j >= start; j-- {
			if s[j] == s[i] {
				// 说明找到一个新的不重复子串 [start, i-1]
				findNewSub = true
				subLength := i - start
				if subLength > maxLength {
					// [start, i-1]
					maxLength = subLength
				}

				start = j + 1 // start 这里可能为 i
				break
			}
		}

		if !findNewSub {
			// [start, i]
			subLength := i + 1 - start
			if subLength > maxLength {
				maxLength = subLength
			}
			// 这里不需要更新 start
		}
	}

	//fmt.Println(left, right, s[left:right+1])

	return maxLength
}

func main() {
	fmt.Println(lengthOfLongestSubstring("ab"))
}

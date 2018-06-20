package main

import (
	"fmt"
	"math"
)

func longestCommonPrefix(strs []string) string {

	num := len(strs)
	if num == 0 {
		return ""
	}

	minLength := math.MaxInt32

	for _, str := range strs {
		length := len(str)
		if length < minLength {
			minLength = length
		}
	}

	current := 0
	for current < minLength {
		elem := strs[0][current]
		stop := false
		for j := 1; j < num; j++ {
			if strs[j][current] != elem {
				stop = true
				break
			}
		}
		if stop {
			break
		}
		current++
	}

	return strs[0][0:current]
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"abc", "abd", "abc"}))
}

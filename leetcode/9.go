package main

import (
	"fmt"
	"math"
)

func isPalindrome(x int) bool {
	if x == 0 {
		return true
	}

	if x < 0 {
		return false
	}

	// 320
	if x%10 == 0 {
		return false
	}

	old := x
	digits := 0
	for x > 0 {
		x = x / 10
		digits++
	}

	// 可以不转化为字符串
	val := 0
	base := int(math.Pow10(digits - 1))
	x = old
	for x > 0 {
		a := x % 10 // 取出最后一位
		x = x / 10

		val += a * base
		base /= 10
	}

	return old == val
}

func main() {
	fmt.Println(isPalindrome(121))
}

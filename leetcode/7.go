package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	isNagative := false

	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}

	if x == 0 {
		return 0
	} else if x < 0 {
		isNagative = true
		x = -x
	}

	digits := make([]int, 0, 10)

	for x > 0 {
		a := x % 10 // 解析出最后一位
		x = x / 10  // 123 变为 12
		digits = append(digits, a)
	}

	// reverse string
	start := 0
	end := len(digits) - 1
	for start < end {
		temp := digits[start]
		digits[start] = digits[end]
		digits[end] = temp
		start++
		end--
	}

	sum := 0
	base := 1
	for _, digit := range digits {
		sum += digit * base
		base *= 10
	}

	if isNagative {
		sum = -sum
	}

	if sum > math.MaxInt32 || sum < math.MinInt32 {
		return 0
	}

	return sum
}

func main() {
	fmt.Println(math.MaxInt32)
	fmt.Println(reverse(1534236469))
}

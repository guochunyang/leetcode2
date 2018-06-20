package main

import (
	"fmt"
	"math"
)

func myAtoi(str string) int {
	length := len(str)

	start := 0
	hasNegative := false
	for start < length {
		// space
		if str[start] <= 32 {
			start++
			continue
		} else if str[start] == '-' {
			hasNegative = true
			start++
			break
		} else if str[start] == '+' {
			start++
			break
		} else if str[start] < '0' || str[start] > '9' {
			return 0
		} else {
			break
		}
	}

	var sum = 0
	for start < length {

		// invalid
		if str[start] < '0' || str[start] > '9' {
			break
		}
		// digit
		dight := str[start] - '0'
		sum = sum*10 + int(dight)

		if sum > math.MaxInt32 {
			break
		}

		start++
	}

	if hasNegative {
		sum = -1 * sum
	}

	if sum > math.MaxInt32 {
		sum = math.MaxInt32
	} else if sum < math.MinInt32 {
		sum = math.MinInt32
	}

	return sum
}

func main() {
	//fmt.Println(myAtoi("    -8974456dfew"))
	//fmt.Println(myAtoi("-91283472332"))
	fmt.Println(myAtoi("9223372036854775808"))
}

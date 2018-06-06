package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	result := make([]int, 0, 2)
	length := len(nums)
	if length == 0 {
		return result
	}

	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				result = append(result, i, j)
			}
		}
	}

	return result
}

func main() {
	nums := []int{3, 3}
	fmt.Println(twoSum(nums, 6))
}

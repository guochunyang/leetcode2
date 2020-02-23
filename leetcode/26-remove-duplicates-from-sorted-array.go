package main

import "fmt"

func removeDuplicates(nums []int) int {

	cur := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[cur] {
			continue
		}
		cur++
		nums[cur] = nums[i]
	}
	return cur + 1
}

func main() {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

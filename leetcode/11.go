package main

import "fmt"

func maxArea(height []int) int {
	length := len(height)

	if length == 0 {
		return 0
	}

	max := 0
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			min := 0
			if height[i] > height[j] {
				min = height[j]
			} else {
				min = height[i]
			}

			area := (j - i) * min
			if area > max {
				max = area
			}
		}
	}

	return max
}

func main() {
	fmt.Println(maxArea([]int{0, 2}))
}

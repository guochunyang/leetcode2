package main

func removeElement(nums []int, val int) int {
	cur := -1 // cur 表示当前坐标位置
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			cur++
			nums[cur] = nums[i]
		}
	}
	return cur + 1
}

func main() {

}

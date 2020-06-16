package main

import "fmt"

func firstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}

	for i := 0; i < len(nums); i++{
		for nums[i] >= 1 && nums[i] < len(nums) && (nums[i] - 1 != i) && (nums[i] != nums[nums[i] - 1]){
			nums[i], nums[nums[i] - 1] = nums[nums[i] - 1], nums[i]
		}
	}

	for i, v := range nums{
		if v - 1 != i {
			return i + 1
		}
	}
	return len(nums) + 1
}


func main() {
	nums := []int{1, 1}
	a := firstMissingPositive(nums)
	fmt.Println(a)

}
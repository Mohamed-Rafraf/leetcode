package main

import "fmt"

func main() {
	array := []int{0, 0, 0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println("Before ", array)
	n := removeDuplicates(array)
	fmt.Println("After  ", array[:n])
}

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	duplicated := 2
	for i := 2; i < len(nums); i++ {
		if nums[duplicated-2] != nums[i] {
			nums[duplicated] = nums[i]
			duplicated++

		}
	}
	return duplicated
}

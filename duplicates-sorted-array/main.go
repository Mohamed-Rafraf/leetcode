package main

import "fmt"

func main() {
	array := []int{0, 0, 0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println("Before ", array)
	n := removeDuplicates(array)
	fmt.Println("After  ", array[:n])
}

func removeDuplicates(nums []int) int {
	duplicated := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			nums[duplicated] = nums[i]
			duplicated++
		}
	}
	return duplicated
}

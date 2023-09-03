package main

import "fmt"

func main() {
	array := []int{1, 1, 2, 2, 2, 1, 1}
	fmt.Println(array)
	fmt.Println("The candidate ", majorityElement(array))
}

func majorityElement(nums []int) int {
	var candidate int
	vote := 0
	for _, element := range nums {
		if vote == 0 {
			candidate = element
		}
		if candidate == element {
			vote++
		} else {
			vote--
		}
	}
	return candidate
}

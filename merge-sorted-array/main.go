package main

import "fmt"

func main() {
	array1 := []int{1, 2, 3, 0, 0, 0}
	array2 := []int{2, 5, 6}
	fmt.Println("1st array: ", array1)
	fmt.Println("2nd array: ", array2)
	merge(array1, 3, array2, 3)
	fmt.Println("result   : ", array1)

}

func merge(nums1 []int, m int, nums2 []int, n int) {

	for n != 0 {
		if m != 0 && nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
}

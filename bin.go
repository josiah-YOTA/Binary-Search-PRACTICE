package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	return binarySearch(0, len(nums)-1, nums, target)
}

func binarySearch(start, end int, nums []int, target int) int {
	if end >= start {
		mid := start + (end-start)/2

		if target == nums[mid] {
			return mid // Target found
		} else if target > nums[mid] {
			return binarySearch(mid+1, end, nums, target)
		}

		return binarySearch(start, mid-1, nums, target)
	}

	return -1 // Target not found
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 3

	result := search(nums, target)
	if result != -1 {
		fmt.Printf("Element found at index %d\n", result)
	} else {
		fmt.Println("Element not found")
	}
}

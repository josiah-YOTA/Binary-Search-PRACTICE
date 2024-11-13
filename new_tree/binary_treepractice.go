package main

import "fmt"

func search(nums []int, target int) int {
	return binarySearch(0, len(nums)-1, nums, target)
}

func binarySearch(start, end int, nums []int, target int) int {
	for start <= end {
		mid := start + (end - start) / 2 // Mid here is just the index, Using indexing I can grab the value within the array with nums[mid]

		if target == nums[mid] {
			return mid // target was reached at the mid point
		} else if target > nums[mid] {
			return binarySearch(mid + 1, end, nums, target)
		} else {
			return binarySearch(start, mid - 1, nums, target)
		}
	}

	return -1 // Target not found

}

func main() {
	nums := []int{1,2,3,4,5,6}
	target := 5
	result := search(nums, target)

	if result == -1 {
		fmt.Println("the Target is not in this array")
	} else {
		fmt.Printf("Target was found at index %d\n", result)
	}


}
package main

import "fmt"

func removeDuplicates(nums []int) int {
	length := len(nums)
	i := 0
	for j := 1; j < length; j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}

func main() {
	array := []int{1, 2, 3,3,4}
	length := removeDuplicates(array)
	fmt.Println(length)
}

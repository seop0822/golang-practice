package main

import "fmt"

func rotate(nums []int, k int) {
	length := len(nums)
	for k > 0 {
		last := nums[length-1]
		k--
		for i := length - 1; i > 0; i-- {
			j := i - 1
			nums[i] = nums[j]
		}
		nums[0] = last
	}
	fmt.Println(nums)
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(array, 3)
}

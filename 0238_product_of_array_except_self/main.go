package main

import "fmt"

func main() {
	case1 := []int{1, 2, 3, 4}
	result := productExceptSelf(case1)
	fmt.Println(result)
}

func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))

	prefix, postfix := 1, 1
	for i := 0; i < len(nums); i++ {
		answer[i] = prefix
		prefix *= nums[i]
	}

	for i := len(nums) - 1; i >= 0; i-- {
		answer[i] *= postfix
		postfix *= nums[i]
	}

	return answer
}

package main

import "fmt"

func main() {
	case1 := []int{1, 2, 3, 1, 1, 3}
	result := numIdenticalPairs(case1)
	fmt.Println(result)
}

func numIdenticalPairs(nums []int) int {
	// write code here
	total := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				total++
			}
		}
	}
	return total
}

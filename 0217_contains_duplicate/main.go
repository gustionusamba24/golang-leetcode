package main

import "fmt"

func main() {
	example1 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	result := containsDuplicate(example1)
	fmt.Println(result)
}

func containsDuplicate(nums []int) bool {
	temp := make(map[int]bool, len(nums))

	for _, v := range nums {
		found := temp[v]
		if found {
			return true
		}
		temp[v] = true
	}
	return false
}

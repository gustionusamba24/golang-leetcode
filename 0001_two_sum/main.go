package main

import "fmt"

func main() {
	nums := []int{1, 3, 7, 9, 2}
	target := 16
	result := twoSum(nums, target)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for i, num := range nums {
		numberToFind := target - num
		if hashMapIndex, ok := hashMap[numberToFind]; ok {
			return []int{hashMapIndex, i}
		}
		hashMap[num] = i
	}

	return []int{}
}

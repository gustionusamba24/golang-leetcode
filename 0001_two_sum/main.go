package main

import "fmt"

func main() {
	nums := []int{1, 3, 7, 9, 2}
	target := 16
	result := twoSum(nums, target)
	fmt.Println(result)
	target2 := 5
	result2 := twoSumBruteForce(nums, target2)
	fmt.Println(result2)
}

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for i, num := range nums {
		numberToFind := target - num
		hashMapIndex, ok := hashMap[numberToFind]
		if ok {
			return []int{hashMapIndex, i}
		}
		hashMap[num] = i
	}

	return []int{}
}

func twoSumBruteForce(nums []int, target int) []int {
	// iterating over both slices
	// checking for the target

	// O(n2) Brute force solution
	for i, left := range nums {
		for j, right := range nums {
			if left+right == target && i != j {
				return []int{i, j}
			}
		}
	}

	return nil
}

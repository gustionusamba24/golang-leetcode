package main

import "fmt"

func main() {
	input := []int{3, 0, 1}
	result := missingNumber(input)
	fmt.Println(result)
}

func missingNumber(nums []int) int {
	arraySum := 0
	for _, num := range nums {
		arraySum += num // 4
	}
	return (len(nums) * (len(nums) + 1) / 2) - arraySum // (3 * 4 / 2) - 4
}

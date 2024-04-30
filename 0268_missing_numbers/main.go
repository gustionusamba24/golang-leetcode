package main

import "fmt"

func main() {
	input := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	result := missingNumber(input)
	fmt.Println(result)
}

func missingNumber(nums []int) int {
	arraySum := 0
	for _, num := range nums {
		arraySum += num // 4
	}
	return (len(nums) * (len(nums) + 1) / 2) - arraySum // (9 * 10 / 2) - 37
}

package main

import "fmt"

func main() {
	case1 := []int{1, 1, 1, 2, 2, 3}
	result := topKFrequent(case1, 2)
	fmt.Println(result)
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)

	for _, num := range nums {
		m[num]++
	}

	hashMap := make(map[int][]int)

	for v, c := range m {
		hashMap[c] = append(hashMap[c], v)
	}

	result := make([]int, 0)
	for i := len(nums); i >= 1; i-- {
		result = append(result, hashMap[i]...)
		if len(result) > k {
			break
		}
	}
	return result[:k]
}

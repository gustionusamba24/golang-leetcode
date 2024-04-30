package main

import (
	"fmt"
	"sort"
)

func main() {
	bagCapacity := []int{2, 3, 4, 5}
	rocks := []int{1, 2, 4, 4}
	additionalRocks := 2
	result := maximumBags(bagCapacity, rocks, additionalRocks)
	fmt.Println(result)
}

func maximumBags(bagCapacity []int, rocks []int, additionalRocks int) int {
	// write code here
	moreRocks, total := make([]int, len(bagCapacity)), 0
	for i := 0; i < len(bagCapacity); i++ {
		moreRocks[i] = bagCapacity[i] - rocks[i]
		total += moreRocks[i]
	}
	if total < additionalRocks {
		return len(rocks)
	}
	sort.Ints(moreRocks)
	count := 0
	for i := 0; i < len(moreRocks) && moreRocks[i] <= additionalRocks; i++ {
		additionalRocks = additionalRocks - moreRocks[i]
		count++
	}
	return count
}

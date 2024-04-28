package main

import (
	"container/heap"
	"fmt"
)

func main() {
	case1 := []int{1, 1, 1, 2, 2, 3}
	result := topKFrequent(case1, 2)
	fmt.Println(result)
}

// func topKFrequent(nums []int, k int) []int {
// 	m := make(map[int]int)

// 	for _, num := range nums {
// 		m[num]++
// 	}

// 	hashMap := make(map[int][]int)

// 	for v, c := range m {
// 		hashMap[c] = append(hashMap[c], v)
// 	}

// 	result := make([]int, 0)
// 	for i := len(nums); i >= 1; i-- {
// 		result = append(result, hashMap[i]...)
// 		if len(result) > k {
// 			break
// 		}
// 	}
// 	return result[:k]
// }

func topKFrequent(nums []int, k int) []int {
	// Initialize the result slice and count map to count frequency of each element
	var result []int
	counterMap := make(map[int]int)

	// frequency count of each element
	for _, v := range nums {
		counterMap[v]++
	}

	// Initialize a heap of elements with their counts
	h := &Heap{}

	for key, v := range counterMap {
		e := element{
			val: key,
			cnt: v,
		}

		// Pushing each element and its frequency into a heap
		heap.Push(h, e)

		// If heap size goes beyond k, pooping out the latest frequent one
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	// Pooping remaining elements in heap and appending them to the result
	for h.Len() > 0 {
		result = append(result, heap.Pop(h).(element).val)
	}

	// Return the result slice when k most frequent elements from nums
	return result
}

// Defining struct to hold element and its count
type element struct {
	val int
	cnt int
}

// Defining Heap as slice of elements
type Heap []element

// Function to return len of heap
func (h Heap) Len() int {
	return len(h)
}

// Function to swap elements at index i and j
func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Function returning true if frequency of element at i is less than at j
func (h Heap) Less(i, j int) bool {
	return h[i].cnt < h[j].cnt
}

// Function to push an element into heap
func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(element))
}

// Function to pop out an element from heap
func (h *Heap) Pop() (x interface{}) {
	*h, x = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

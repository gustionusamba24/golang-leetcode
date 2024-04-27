package main

import "fmt"

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}

func groupAnagrams(strs []string) [][]string {
	hashMap := make(map[[26]int][]string, 0)

	for _, str := range strs {
		var key [26]int
		for _, char := range str {
			key[char-'a']++
		}
		hashMap[key] = append(hashMap[key], str)
	}

	temp := make([][]string, 0)
	for _, val := range hashMap {
		temp = append(temp, val)
	}
	return temp
}

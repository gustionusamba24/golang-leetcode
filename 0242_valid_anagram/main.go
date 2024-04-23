package main

import "fmt"

func main() {
	result := isAnagram("anagram", "nagaram")
	fmt.Println(result)
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var sAlphabetCounter [26]int
	var tAlphabetCounter [26]int

	for i := 0; i < len(s); i++ {
		sAlphabetCounter[s[i]-'a']++
		tAlphabetCounter[t[i]-'a']++
	}

	for j := 0; j < len(sAlphabetCounter); j++ {
		if sAlphabetCounter[j] != tAlphabetCounter[j] {
			return false
		}
	}

	return true
}

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

	var alphabetCounter [26]int

	for i := 0; i < len(s); i++ {
		alphabetCounter[s[i]-'a']++
		alphabetCounter[t[i]-'a']--
	}

	for j := 0; j < len(alphabetCounter); j++ {
		if alphabetCounter[j] != 0 {
			return false
		}
	}

	return true
}

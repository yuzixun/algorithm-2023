package main

func canConstruct(ransomNote string, magazine string) bool {
	mm := make(map[byte]int)

	for i := 0; i < len(magazine); i++ {
		mm[magazine[i]]++
	}

	for i := 0; i < len(ransomNote); i++ {
		mm[ransomNote[i]]--
		if mm[ransomNote[i]] < 0 {
			return false
		}
	}
	return true
}

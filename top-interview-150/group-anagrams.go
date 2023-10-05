package main

func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)

	for _, str := range strs {
		byteCount := [26]int{}
		for _, b := range str {
			byteCount[b-'a']++
		}
		m[byteCount] = append(m[byteCount], str)
	}

	var result [][]string
	for _, strs := range m {
		result = append(result, strs)
	}
	return result
}

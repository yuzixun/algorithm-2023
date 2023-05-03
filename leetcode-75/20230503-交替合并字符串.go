package leetcode_75

func mergeAlternately(word1 string, word2 string) string {
	var (
		i, index1, index2 = 0, 0, 0
		res               = make([]byte, len(word1)+len(word2))
	)

	for ; index1 < len(word1) && index2 < len(word2); i++ {
		if i%2 == 0 {
			res[i] = word1[index1]
			index1++
		} else {
			res[i] = word2[index2]
			index2++
		}
	}

	for ; index1 < len(word1); i++ {
		res[i] = word1[index1]
		index1++
	}
	for ; index2 < len(word2); i++ {
		res[i] = word2[index2]
		index2++
	}
	return string(res)
}

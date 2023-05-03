package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	str := strs[0]
	size := len(str)
	for i := 1; i < len(strs); i++ {
		minSize := min(size, len(strs[i]))
		j := 0
		for ; j < minSize; j++ {
			if str[j] != strs[i][j] {
				break
			}
		}
		str = str[:j]
		size = len(str)
		if size == 0 {
			break
		}
	}
	return str
}

package leetcode_75

func isSubsequence(s string, t string) bool {
	sPointer, tPointer := 0, 0

	for sPointer < len(s) && tPointer < len(t) {
		if s[sPointer] == t[tPointer] {
			sPointer++
			tPointer++
		}
		tPointer++
	}

	return sPointer == len(s)-1
}

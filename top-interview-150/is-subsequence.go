package main

func isSubsequence(s string, t string) bool {
	sPointer, tPointer := 0, 0

	for ; sPointer < len(s) && tPointer < len(t); tPointer++ {
		if s[sPointer] == t[tPointer] {
			sPointer++
			continue
		}
	}

	return sPointer == len(s)
}

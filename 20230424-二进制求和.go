package main

import "strconv"

func addBinary(a string, b string) string {
	var (
		res        string
		carry      int
		aLen, bLen = len(a), len(b)
		maxLen     = max(aLen, bLen)
	)

	for i := 0; i < maxLen; i++ {
		if i < aLen {
			carry += int(a[i] - '0')
		}
		if i < bLen {
			carry += int(b[i] - '0')
		}

		res = strconv.Itoa(carry%2) + res
		carry = carry / 2
	}
	if carry > 0 {
		res = "1" + res
	}

	return res
}

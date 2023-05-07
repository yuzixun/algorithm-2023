package leetcode_75

import "fmt"

func compress(chars []byte) int {
	write, anchor := 0, 0
	for read := 0; read < len(chars); read++ {
		if read == len(chars)-1 || chars[read+1] != chars[read] {
			chars[write] = chars[anchor]
			write++
			if read > anchor {
				count := read + 1 - anchor
				for _, b := range []byte(fmt.Sprint(count)) {
					chars[write] = b
					write++
				}
			}
			anchor = read + 1
		}
	}
	return write
}

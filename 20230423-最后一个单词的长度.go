package main

//func lengthOfLastWord(s string) int {
//	var (
//		res   = 0
//		index = len(s) - 1
//	)
//	for s[index] == ' ' {
//		index--
//	}
//	for index >= 0 && s[index] != ' ' {
//		index--
//		res++
//	}
//
//	return res
//}

func lengthOfLastWord(s string) int {
	tail := len(s) - 1
	for tail >= 0 && s[tail] == ' ' {
		tail--
	}
	if tail < 0 {
		return 0
	}
	head := tail
	for head >= 0 && s[head] != ' ' {
		head--
	}
	return tail - head
}

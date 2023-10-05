package main

func isValid(s string) bool {
	symbols := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	var arr []byte
	for i := 0; i < len(s); i++ {
		pre, exist := symbols[s[i]]
		if !exist {
			arr = append(arr, s[i])
			continue
		}
		if len(arr) == 0 || arr[len(arr)-1] != pre {
			return false
		}
		arr = arr[:len(arr)-1]
	}

	return len(arr) == 0
}

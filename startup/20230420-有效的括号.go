package main

var (
	m = map[rune]rune{
		'}': '{', ']': '[', ')': '(',
	}
)

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	var stack []rune
	for _, i := range s {
		if m[i] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != m[i] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, i)
		}
	}
	return len(stack) == 0
}

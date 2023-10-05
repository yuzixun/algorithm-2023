package main

func isHappy(n int) bool {
	m := make(map[int]int)

	for {
		_, exist := m[n]
		if exist {
			return false
		}
		next := transform(n)
		if next == 1 {
			return true
		}
		m[n] = next
		n = next
	}
}

func transform(n int) int {
	sum := 0
	for n != 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}

package main

//var cache = make(map[int]int)
//
//func climbStairs(n int) int {
//	if n <= 2 {
//		return n
//	}
//
//	if v, ok := cache[n]; ok {
//		return v
//	}
//
//	cache[n] = climbStairs(n-1) + climbStairs(n-2)
//
//	return cache[n]
//}

func climbStairs(n int) int {
	a, b, c := 0, 0, 1
	for i := 0; i < n; i++ {
		a = b
		b = c
		c = a + b
	}
	return c
}

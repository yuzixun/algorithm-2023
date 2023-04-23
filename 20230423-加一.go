package main

//func plusOne(digits []int) []int {
//	var take int
//	for i := len(digits) - 1; i >= 0; i-- {
//		cur := digits[i] + take
//		if i == len(digits)-1 {
//			cur++
//		}
//		digits[i], take = cur%10, cur/10
//	}
//
//	if take > 0 {
//		res := make([]int, len(digits)+1)
//		res = []int{1}
//		res = append(res, digits...)
//		return res
//	}
//	return digits
//}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	res := make([]int, len(digits)+1)
	res = []int{1}
	res = append(res, digits...)
	return res
}

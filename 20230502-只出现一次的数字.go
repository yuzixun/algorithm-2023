package main

func singleNumber(nums []int) int {
	var v int
	for _, number := range nums {
		v = v ^ number
	}
	return v
}

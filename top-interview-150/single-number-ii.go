package main

func singleNumber(nums []int) int {

	res := int32(0)

	for i := 0; i < 32; i++ {
		cur := int32(0)
		for _, num := range nums {
			cur += (int32(num) >> i) & 1
		}

		if cur%3 > 0 {
			res |= 1 << i
		}

	}

	return int(res)
}

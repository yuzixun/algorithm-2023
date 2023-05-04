package leetcode_75

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i, size := 0, len(flowerbed); i < size && n > 0; {
		if flowerbed[i] == 1 {
			i += 2
		} else if i == size-1 || flowerbed[i+1] == 0 {
			n--
			i += 2
		} else {
			i += 3
		}
	}
	return n <= 0
}

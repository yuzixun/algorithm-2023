package leetcode_75

func kidsWithCandies(candies []int, extraCandies int) []bool {
	res := make([]bool, len(candies))
	maxCandies := -1
	for _, candie := range candies {
		maxCandies = max(maxCandies, candie)
	}

	for i, candie := range candies {
		if candie+extraCandies >= maxCandies {
			res[i] = true
		}
	}

	return res
}

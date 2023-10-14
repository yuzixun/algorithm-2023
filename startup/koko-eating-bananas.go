package main

import "math"

func minEatingSpeed(piles []int, h int) int {

	maxPiles := math.MinInt64

	for i := 0; i < len(piles); i++ {
		maxPiles = max(maxPiles, piles[i])
	}

	left, right := 0, maxPiles

	for left <= right {
		mid := left + (right-left)/2
		curHour := minEatingSpeedCalc(piles, mid)
		switch {
		case curHour == h:
			right = mid - 1
		case curHour < h:
			left = mid + 1
		case curHour > h:
			right = mid - 1
		}
	}
	return left - 1
}

func minEatingSpeedCalc(piles []int, speed int) int {
	hour := 0
	for i := 0; i < len(piles); i++ {
		hour += piles[i] / speed
		if piles[i]%speed > 0 {
			hour++
		}
	}
	return hour
}

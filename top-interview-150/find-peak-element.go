package main

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2
		switch {
		case nums[mid] >= nums[mid+1]:
			right = mid
		case nums[mid] < nums[mid+1]:
			left = mid + 1
		}
	}

	// fmt.Println(left, right)
	return left
}

//func findPeakElement(nums []int) int {
//    if len(nums) == 1{
//			return 0
//		}
//	left, right := 0, len(nums)-1
//
//	for left <= right {
//		mid := left + (right-left)/2
//		// fmt.Println(left, right,mid, nums[mid])
//		if mid +1 > len(nums) -1 {
//			break
//		}
//		switch {
//		case nums[mid] >= nums[mid+1]:
//			right = mid -1
//		case nums[mid] < nums[mid+1]:
//			left = mid + 1
//		}
//	}
//
//	// fmt.Println(left, right)
//	return left
//}

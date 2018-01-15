package maximumSubarray

import "math"

//Find the contiguous subarray within an array (containing at least one number) which has the largest sum.
//
//For example, given the array [-2,1,-3,4,-1,2,1,-5,4],
//the contiguous subarray [4,-1,2,1] has the largest sum = 6.

func MaxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sum := nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if sum > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}

		max = int(math.Max(float64(sum), float64(max)))
	}

	return max
}

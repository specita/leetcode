package searchInsertPosition

//Given a sorted array and a target value, return the index if the target is found. If not,
//return the index where it would be if it were inserted in order.
//
//You may assume no duplicates in the array.
//
//Example 1:
//
//Input: [1,3,5,6], 5
//Output: 2
//Example 2:
//
//Input: [1,3,5,6], 2
//Output: 1
//Example 3:
//
//Input: [1,3,5,6], 7
//Output: 4
//Example 1:
//
//Input: [1,3,5,6], 0
//Output: 0

func SearchInsert(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, start, end, target int) int {
	if start > end {
		return end + 1
	}

	mid := (start + end) / 2

	if nums[mid] == target {
		return mid
	} else if nums[mid] > target {
		return binarySearch(nums, start, mid-1, target)
	} else {
		return binarySearch(nums, mid+1, end, target)
	}
}

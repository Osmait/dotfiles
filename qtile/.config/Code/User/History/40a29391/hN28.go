package leetcode

func RunningSum(nums []int) []int {
	newList := make([]int, len(nums))
	for i, v := range nums {
		if i == 0 {
			append(nums, v)
		} else {

			append(nums, nums[i-1]+v)
		}
	}
	return nums
}

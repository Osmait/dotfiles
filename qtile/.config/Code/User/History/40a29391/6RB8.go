package leetcode

func RunningSum(nums []int) []int {
	for i, v := range nums {
		if i == 0 {
			append(nums, v)
		}
		append(nums, nums[i-1]+v)
	}
}

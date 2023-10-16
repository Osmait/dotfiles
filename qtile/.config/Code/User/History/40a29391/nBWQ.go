package leetcode

func RunningSum(nums []int) []int {
	for i, v := range nums {
		if i == 0 {
			append(nums, i)
		}
		nums[i] += nums[i-1]
	}
}

package leetcode



func RunningSum(nums []int) []int {
	for _, i range nums{
		if i == 0{
			append(nums,nums[i])
		}
		nums[i] += nums[i-1]
	}
}
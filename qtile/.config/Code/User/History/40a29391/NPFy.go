package leetcode



func RunningSum(nums []int) []int {
	for i , _ range nums{
		if i == 0{
			append(nums,nums[i])
		}
		nums[i] += nums[i-1]
	}
}
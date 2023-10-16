package leetcode

func RunningSum(nums []int) []int {
	newList := make([]int, len(nums))
	for i, v := range nums {
		if i == 0 {
			newLis = append(newList, v)
		} else {

			newList = append(newList, nums[i-1]+v)
		}
	}
	return newList
}

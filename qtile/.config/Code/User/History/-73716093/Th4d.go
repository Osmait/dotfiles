package leetcode

func PivotIndex(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	leftNum := 0

	for i, v := range nums {
		if leftNum == total-v-leftNum {
			return i
		}
		leftNum += v
	}
	return -1
}

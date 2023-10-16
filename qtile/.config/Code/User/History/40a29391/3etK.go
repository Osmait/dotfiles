package leetcode

func RunningSum(nums []int) []int {
	newList := []int{}
	for _, n := range nums {
		if len(newList) == 0 {
			newList = append(newList, n)
		} else {
			newList = append(newList, n+newList[len(newList)-1])
		}
	}

	return newList

}

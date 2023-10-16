package leetcode

func ContainDuplicate(nums []int) bool {
	hashset := make(map[int]int)
	for _, num := range nums {
		if _, ok := hashset[num]; ok {
			return true
		} else {
			hashset[num] = 1
		}

	}
	return false

}

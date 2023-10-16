package leetcode

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	freq := [26]int{}

	for i := 0; i < len(s); i++ {
		freq[i-'a']++
		freq[i-'a']--
	}
	return
}

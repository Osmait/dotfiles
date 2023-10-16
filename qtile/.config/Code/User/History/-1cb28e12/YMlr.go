package leetcode

import (
	"sort"
)

func IsAnagram(s string, t string) bool {
	// if len(s) != len(t) {
	// 	return false
	// }
	// freq := [26]int{}

	// for i := 0; i < len(s); i++ {
	// 	freq[s[i]-'a']++
	// 	freq[t[i]-'a']--
	// }

	// for _, v := range freq {
	// 	if v != 0 {
	// 		return false
	// 	}
	// }
	// return true

	// return sort.Strings(strings.Split(s, "")) == sort.Strings(strings.Split(t, ""))
	return sort.StringsAreSorted([]string{s, t})
}

package leetcode

func IsISomorphic(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := make(map[byte]byte)
	tMap := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		c1 := s[i]
		c2 := t[i]
		if val, ok := tMap[c2]; ok {
			if val != strings.c2 {

			}
		}

	}

	return true
}

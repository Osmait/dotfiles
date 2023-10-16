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
			if val != c2 {
				return false
			}
		}
		if val, ok := sMap[c1]; ok {
			if val != c1 {
				return false
			}
		}

		sMap[c1] = c2
		tMap[c2] = c1

	}

	return true
}

package lengthoflongestsubstring

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	var (
		res       int
		tmp       int
		existsMap = make(map[uint8]struct{})
		index     int
	)
	for tmpIndex := index; tmpIndex < len(s); {
		if _, ok := existsMap[s[tmpIndex]]; !ok {
			tmp++
			existsMap[s[tmpIndex]] = struct{}{}
			tmpIndex++
			continue
		}
		if res < tmp {
			res = tmp
		}
		tmp = 0
		index++
		tmpIndex = index
		existsMap = make(map[uint8]struct{})
	}
	if res < tmp {
		res = tmp
	}
	return res
}
package longestpalindrome

// https://leetcode.cn/problems/longest-palindromic-substring/
func longestPalindrome(s string) string {
	var res string
	for index := 0; index < len(s); index++ {
		tmp := findSingle(s, index)
		if len(tmp) > len(res) {
			res = tmp
		}
		tmp = findDouble(s, index)
		if len(tmp) > len(res) {
			res = tmp
		}
	}
	return res
}

func findSingle(s string, index int) string {
	var (
		i, j   = index - 1, index + 1
		length = 1
	)
	for {
		if i < 0 {
			break
		}
		if j >= len(s) {
			break
		}
		if s[i] != s[j] {
			break
		}
		length += 2
		i--
		j++
	}
	return s[index-(length/2) : index+(length/2)+1]
}

func findDouble(s string, index int) string {
	var (
		i, j   = index, index + 1
		length int
	)
	for {
		if i < 0 {
			break
		}
		if j >= len(s) {
			break
		}
		if s[i] != s[j] {
			break
		}
		length += 2
		i--
		j++
	}
	return s[index-(length/2)+1 : index+(length/2)+1]
}
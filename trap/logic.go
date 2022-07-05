package trap

// https://leetcode.cn/problems/trapping-rain-water/
func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	min := height[0]
	if height[len(height)-1] < min {
		min = height[len(height)-1]
	}
	var (
		index = 0
		max   = min
	)
	for i, item := range height {
		if i == 0 || i == len(height)-1 {
			continue
		}
		if max >= item {
			continue
		}
		max = item
		index = i
	}
	if index == 0 {
		return sumSubList(height)
	}
	var res int
	res += trap(height[:index+1])
	res += trap(height[index:])
	return res
}

func sumSubList(subList []int) int {
	if len(subList) <= 2 {
		return 0
	}
	min := subList[0]
	if min > subList[len(subList)-1] {
		min = subList[len(subList)-1]
	}
	var res = min * (len(subList) - 2)
	for i := 1; i <= len(subList)-2; i++ {
		res -= subList[i]
	}
	return res
}
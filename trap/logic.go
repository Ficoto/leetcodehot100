package trap

// https://leetcode.cn/problems/trapping-rain-water/
func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	if height[0] >= height[len(height)-1] {
		var (
			index int
			area  int
		)
		for i := len(height) - 2; i >= 0; i-- {
			if height[i] < height[len(height)-1] {
				area += height[len(height)-1] - height[i]
				continue
			}
			index = i
			break
		}
		if index == 0 {
			return area
		}
		return area + trap(height[:index+1])
	}
	var (
		index int
		area  int
	)
	for i := 1; i < len(height); i++ {
		if height[0] > height[i] {
			area += height[0] - height[i]
			continue
		}
		index = i
		break
	}
	if index == len(height)-1 {
		return area
	}
	return area + trap(height[index:])
}

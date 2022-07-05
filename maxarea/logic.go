package maxarea

// https://leetcode.cn/problems/container-with-most-water/submissions/
func maxArea(height []int) int {
	var (
		leftIndex  int
		rightIndex = len(height) - 1
		maxRes     int
	)
	for {
		if leftIndex >= rightIndex {
			break
		}
		h := height[leftIndex]
		if height[leftIndex] >= height[rightIndex] {
			h = height[rightIndex]
			rightIndex--
		} else {
			leftIndex++
		}
		tmp := (rightIndex - leftIndex + 1) * h
		if tmp > maxRes {
			maxRes = tmp
		}
	}
	return maxRes
}
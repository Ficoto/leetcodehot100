package twosum

// https://leetcode.cn/problems/two-sum/
func twoSum(nums []int, target int) []int {
	var (
		i, j = 0, 1
	)
	for {
		if i >= len(nums) {
			break
		}
		if j >= len(nums) {
			i++
			j = i + 1
		}
		if nums[i]+nums[j] != target {
			j++
			continue
		}
		break
	}
	return []int{i, j}
}
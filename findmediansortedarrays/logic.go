package findmediansortedarrays

// https://leetcode.cn/problems/median-of-two-sorted-arrays/
func findMedianSortedArrays(num1 []int, num2 []int) float64 {
	length := len(num1) + len(num2)
	isDouble := length%2 == 0
	var k = length / 2
	if !isDouble {
		k += 1
	}
	return findKIndexNumber(num1, num2, k, isDouble)
}

func findKIndexNumber(num1, num2 []int, k int, isDouble bool) float64 {
	if len(num1) == 0 && len(num2) == 0 {
		return 0
	}
	if len(num1) == 0 {
		if !isDouble {
			return float64(num2[k-1])
		}
		return float64(num2[k-1]+num2[k]) / 2
	}
	if len(num2) == 0 {
		if !isDouble {
			return float64(num1[k-1])
		}
		return float64(num1[k-1]+num1[k]) / 2
	}
	if k <= 1 {
		if !isDouble {
			if num1[0] < num2[0] {
				return float64(num1[0])
			}
			return float64(num2[0])
		}
		var target1, target2 int
		if num1[0] < num2[0] {
			target1 = num1[0]
			if len(num1) > 1 && num1[1] < num2[0] {
				target2 = num1[1]
			} else {
				target2 = num2[0]
			}
		} else {
			target1 = num2[0]
			if len(num2) > 1 && num2[1] < num1[0] {
				target2 = num2[1]
			} else {
				target2 = num1[0]
			}
		}
		return float64(target1+target2) / 2
	}
	var num1Index, num2Index = k/2 - 1, k/2 - 1
	if len(num1)-1 < num1Index {
		num1Index = len(num1) - 1
	}
	if len(num2)-1 < num2Index {
		num2Index = len(num2) - 1
	}
	if num1[num1Index] < num2[num2Index] {
		return findKIndexNumber(num1[num1Index+1:], num2, k-num1Index-1, isDouble)
	}
	return findKIndexNumber(num1, num2[num2Index+1:], k-num2Index-1, isDouble)
}
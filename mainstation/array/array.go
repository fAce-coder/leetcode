package array

import (
	"math"
)

// maxOfIntArray
//
//	@Description:求int类型数组中的最大值
func maxOfIntArray(nums []int) int {
	maxOfArray := math.MinInt
	for _, num := range nums {
		if num > maxOfArray {
			maxOfArray = num
		}
	}
	return maxOfArray
}

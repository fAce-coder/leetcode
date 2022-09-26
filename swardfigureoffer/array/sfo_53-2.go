package array

func missingNumber(nums []int) int {
	// 二分法

	// 左右指针
	var left int = 0
	var right int = len(nums) - 1

	// 二分查找
	for left <= right {
		mid := (left + right) / 2
		// 值与索引相同，说明缺失值在右边
		if nums[mid] == mid {
			left = mid + 1
			// 值大于索引值，说明缺失值在左边
		} else if nums[mid] > mid {
			right = mid - 1
		}
	}
	// 最后left所在位置的值就是缺失值的下一个值，所在的索引就是缺失值
	return left
}

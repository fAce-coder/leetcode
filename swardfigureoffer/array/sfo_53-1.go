package array

func search(nums []int, target int) int {
	// 二分法
	// 找到右边界，找到左边界，右边界-左边界 就是该数出现的次数

	// 特殊情况
	if len(nums) == 0 {
		return 0
	}
	// 左右指针
	var left int = 0
	var right int = len(nums) - 1
	// 左右边界
	var L int
	var R int
	// 找右边界
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			left = mid + 1 // 因为要找右边界，所以如果相等的情况，左指针还是要右移
		}
	}
	R = right // 跳出循环时，此时right就是右边界
	// 判断此时右边界是否是要找的target，如果不是，说明数组中没有目标数target,函数结束
	if R == -1 || nums[R] != target {
		return 0
	}
	// 找左边界
	// 右指针不用重置，可以减小时间复杂度
	left = 0
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1 // 因为要找左边界，所以如果相等的情况，右指针还是要左移
		}
	}
	L = left // 跳出循环时，此时left就是左边界
	// 返回结果
	return R - L + 1
}

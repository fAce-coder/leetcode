package array

func searchRange(nums []int, target int) []int {
	// 二分查找

	/*
		条件：有序数组、寻找目标值、时间复杂度logn，完全符合二分查找，首先使用二分查找方法
		因为目标数可能不止是一个，可能是多个，因此要返回该目标数的索引范围，因此考虑找到该目标数的左右边界
		先二分查找右边界：
			如果mid处值小于目标值，则left=mid+1
			如果mid处值大于目标值，则right=mid-1
			如果mid处值等于目标值，则为了继续向右找右边界，left=mid+1
			直到left和right越界，此时right位置就是该数的右边界
		再二分查找左边界：
			如果mid处值小于目标值，则left=mid+1
			如果mid处值大于目标值，则right=mid-1
			如果mid处值等于目标值，则为了继续向左找左边界，right=mid-1
			直到left和right越界，此时left位置就是该数的左边界
	*/

	// 1.特殊情况，如果数组为空，则返回[-1,-1]
	n := len(nums)
	if n == 0 {
		return []int{-1, -1}
	}

	// 2.初始化
	// 2.1 左右边界
	L := 0
	R := 0
	// 2.2 左右指针
	left := 0
	right := n - 1

	// 3.查找右边界
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			// 3.1 如果mid处值小于目标值，则left=mid+1
			left = mid + 1
		} else if nums[mid] > target {
			// 3.2 如果mid处值大于目标值，则right=mid-1
			right = mid - 1
		} else {
			// 3.3 如果mid处值等于目标值，则为了继续向右找右边界，left=mid+1
			left = mid + 1
		}
	}
	// 3.4 直到left和right越界，此时right位置就是该数的右边界
	R = right
	// 3.5 如果右边界越界，或者右边界的数不是目标数，说明目标数不在数组中，此时不用继续判断左边界，直接返回[-1,-1]
	if R < 0 || nums[R] != target {
		return []int{-1, -1}
	}

	// 4.查找左边界
	// 将左边界归位，右边界不需要归位（右边界此时就是目标数的右边界，无需归位，减少查找次数和时间消耗）
	left = 0
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			// 4.1 如果mid处值小于目标值，则left=mid+1
			left = mid + 1
		} else if nums[mid] > target {
			// 4.2 如果mid处值大于目标值，则right=mid-1
			right = mid - 1
		} else {
			// 4.3 如果mid处值等于目标值，则为了继续向左找左边界，right=mid-1
			right = mid - 1
		}
	}
	// 4.4 直到left和right越界，此时left位置就是该数的左边界
	L = left

	// 5.返回目标值的左右边界
	return []int{L, R}
}

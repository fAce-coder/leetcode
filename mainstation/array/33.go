package array

func search(nums []int, target int) int {
	// 二分查找

	/*
		条件：数组有序（两个有序数组段）、查找某个数、时间复杂度要求nlogn；完全符合二分查找，因此首先考虑二分查找
		因为数组不是整体有序，而是发生了旋转，因此变成了两个有序数组段，因此需要先判断target在哪一段有序数组段中，再在这个有序数组段中进行二分查找
		在左右索引不越界的情况下：
			1.如果nums[mid]==target，说明找到了，则直接返回结果
			2.如果没找到，则尝试为target找到所在的有序数组段
				2.1 如果nums[mid]>nums[right]，说明mid在左半边有序数组段内，即left到mid这一段是有序的；此时判断target是否在left到mid这一段有序数组段中，是则在其中二分查找，不是则在mid到right中继续找其所属的有序数组段
				2.2 如果nums[mid]<nums[right]，说明mid在右半边有序数组段内，即mid到right这一段是有序的；此时判断target是否在mid到right这一段有序数组段中，是则在其中二分查找，不是则在left到mid中继续找其所属的有序数组段
	*/

	// 1.初始化左右指针
	left := 0
	right := len(nums) - 1

	// 2.在左右索引不越界的情况下，进行二分查找
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			// 2.1 如果nums[mid]==target，说明找到了，则直接返回结果
			return mid
		} else if nums[mid] > nums[right] {
			// 2.2 如果nums[mid]>nums[right]，说明mid在左半边有序数组段内，即left到mid这一段是有序的
			// 此时判断target是否在left到mid这一段有序数组段中，是则在其中二分查找；不是则在mid到right中继续找其所属的有序数组段
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// 2.3 如果nums[mid]<nums[right]，说明mid在右半边有序数组段内，即mid到right这一段是有序的
			// 此时判断target是否在mid到right这一段有序数组段中，是则在其中二分查找，不是则在left到mid中继续找其所属的有序数组段
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	// 3.如果左右指针越界还没找到，则说明target不在数组内，返回-1
	return -1
}

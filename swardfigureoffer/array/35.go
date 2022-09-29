package array

func searchInsert(nums []int, target int) int {
	// 二分查找

	/*
		有序数组、数组中无重复数值、要求时间复杂度是log n
		完全符合二分查找的要求
	*/

	// 1.初始化左右索引
	left := 0
	right := len(nums) - 1

	// 2.当左右未越界
	for left <= right {
		// 2.1 找出中间节点
		mid := (left + right) / 2
		// 2.2 比较目标值和中间节点值的大小
		if target > nums[mid] {
			// 2.2.1 如果目标值大于中间节点值，则将左索引移到中间节点右侧
			left = mid + 1
		} else if target < nums[mid] {
			// 2.2.2 如果目标值小于中间节点值，则将右索引移到中间节点左侧
			right = mid - 1
		} else {
			// 2.2.3 如果目标值等于中间节点值，则说明找到该节点，返回该节点索引
			return mid
		}
	}

	// 3.如果左右越界还没有找到目标节点，说明切片中不存在该节点，此时left就是其应该插入的位置
	return left
}

package array

func twoSum1(nums []int, target int) []int {
	// 哈希表

	var res []int = make([]int, 2)
	// 新建哈希表
	var hashTable map[int]bool = make(map[int]bool)

	for _, num := range nums {
		if len(hashTable) == 0 || !hashTable[target-num] {
			hashTable[num] = true
		} else {
			res[0] = num
			res[1] = target - num
			break
		}
	}

	return res
}

func twoSum(nums []int, target int) []int {
	// 头尾双指针

	// 头尾指针定义
	var left int = 0
	var right int = len(nums) - 1

	// 返回结果定义
	var res []int = make([]int, 2)

	for left < right {
		sumOfLeftAndRight := nums[left] + nums[right]
		// 头尾指针的和 大于 目标值，右指针左移
		if sumOfLeftAndRight > target {
			right--
			// 头尾指针的和 小于 目标值，左指针右移
		} else if sumOfLeftAndRight < target {
			left++
			// 头尾指针的和 等于 目标值，返回
		} else {
			res[0] = nums[left]
			res[1] = nums[right]
			break
		}
	}

	return res
}

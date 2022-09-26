package array

func exchange1(nums []int) []int {
	// 头尾双指针

	// 左右边界
	var left int = 0
	var right int = len(nums) - 1

	for left < right {
		// 从左面开始找到1个偶数
		for left < right && nums[left]%2 != 0 {
			left++
		}
		// 从右边开始找到1个奇数
		for left < right && nums[right]%2 == 0 {
			right--
		}
		// 交换这两个数
		nums[left], nums[right] = nums[right], nums[left]
	}

	return nums
}

func exchange2(nums []int) []int {
	// 快慢双指针
	// 快指针找奇数，来跟慢指针交换

	// 指针定义
	var slow int = 0
	var fast int = 0

	var rightOfNums int = len(nums) - 1

	for fast <= rightOfNums {
		// 快指针找到奇数，与满指针交换
		if nums[fast]%2 != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			fast++
			slow++
			// 快指针没有找到奇数，就继续往后找
		} else {
			fast++
		}
	}

	return nums
}

func exchange3(nums []int) []int {
	// 双数组
	// 一个数组存奇数，一个数组存偶数，最后将两个数组合并

	var single []int = make([]int, 0)
	var double []int = make([]int, 0)

	for _, num := range nums {
		// 奇数，放到single数组中
		if num%2 != 0 {
			single = append(single, num)
			// 偶数，放到double数组中
		} else {
			double = append(double, num)
		}
	}
	// 将两个数组合并
	var res []int = append(single, double...)

	return res
}

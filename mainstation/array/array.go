package array

import (
	"math"
)

// 可排序切片
type sortableNums []int

func (n sortableNums) Len() int {
	return len(n)
}

func (n sortableNums) Less(i, j int) bool {
	// 从小到大排序
	return n[i] < n[j]
}

func (n sortableNums) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

// 可排序二维数组
type sortable2DNums [][]int

func (s sortable2DNums) Len() int {
	return len(s)
}

func (s sortable2DNums) Less(i, j int) bool {
	if s[i][0] < s[j][0] {
		// 按照数组中第一个数从小到大排序
		return true
	} else if s[i][0] == s[j][0] {
		// 如果数组中第一个数相同，则按照数组中第二个数从小到大排序
		return s[i][1] < s[j][1]
	} else {
		return false
	}
}

func (s sortable2DNums) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// 快速排序
func quickSort(nums []int) {
	// 1.定义两个匿名函数用来递归
	var partition func([]int, int, int) int
	var quicksort func([]int, int, int)

	// 2.分区函数
	partition = func(nums []int, left int, right int) int {
		// 2.1 让基准为最左侧元素
		key := left
		// 2.2 在当前基准下进行交换
		for left < right {
			// 2.2.1 在右侧找到第一个小于基准的数
			for left < right && nums[right] >= nums[key] {
				right--
			}
			// 2.2.2 在左侧找到第一个大于基准的数
			for left < right && nums[left] <= nums[key] {
				left++
			}
			// 2.2.3 将这两个数交换,使得前面的都是小于key的数,后面的都是大于key的数
			nums[left], nums[right] = nums[right], nums[left]
		}
		// 2.3 本轮交换完毕后，将基准放到当前left处，即该基准数应该在的位置
		nums[key], nums[left] = nums[left], nums[key]
		// 2.4 返回当前基准的索引
		return left
	}

	// 3.递归排序函数
	quicksort = func(nums []int, left int, right int) {
		// 3.1 越界则返回
		if left >= right {
			return
		}
		// 3.2 当前递归中要做的事：进行一轮排序，并将基准数的新位置索引返回
		mid := partition(nums, left, right)
		// 3.3 递归：以该基准数的新位置索引为中点，递归地对左边和右边的无序数组排序
		quicksort(nums, left, mid-1)
		quicksort(nums, mid+1, right)
	}

	// 4.主函数
	// 4.1 如果数组长度<=1，则不用排序直接返回
	lengthOfNums := len(nums)
	if lengthOfNums <= 1 {
		return
	}
	// 4.2 进行快速排序
	left := 0
	right := lengthOfNums - 1
	quicksort(nums, left, right)
}

// 求int类型的绝对值
func absInt(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}

// 求int数组中的最大值
func maxInt(nums ...int) int {
	maxOfNums := math.MinInt
	for _, num := range nums {
		if num > maxOfNums {
			maxOfNums = num
		}
	}
	return maxOfNums
}

// 求int数组中的最小值
func minInt(nums ...int) int {
	minOfNums := math.MaxInt
	for _, num := range nums {
		if num < minOfNums {
			minOfNums = num
		}
	}
	return minOfNums
}

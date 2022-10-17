package array

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	// 排序+三指针

	/*
		先将数组从小到大排序
		三指针思路：先固定指针k，然后找合适的i和j
		指针k：
			指针k从数组中最小的元素开始固定，固定下之后就开始找对应的i、j
			每轮找完之后指针k右移一位，直到指针k所在的值>0（因为k是最小的，k都大于0了，则i、j也一定大于0，因此三者的和不可能等于0）或者指针k索引超过n-2（因为要求i!=j!=k，因此如果k超过n-2，则k、i、j必定会有两个相等的索引，不符合条件），此时退出函数
		指针i、j：
			指针i放在k之后的一位，指针j放在数组的最后
			如果i、j、k三者和小于0，则指针i右移，寻找更大元素
			如果i、j、k三者和大于0，则指针j左移，寻找更小元素
		因为要求不能有重复的组合，因此三个指针遇到重复的数时，都继续移动去重，直到不再重复
	*/

	// 1.初始化结果集
	res := make([][]int, 0)

	// 2.调用标准库，对数组进行从小到大排序
	sort.Sort(sortableNums(nums))

	// 3.指针k从数组中最小的元素开始固定，逐渐右移
	n := len(nums)
	for k := 0; k < n-2; k++ {
		// 3.1 如果指针k所在的数>0，则退出
		// 因为k是最小的，k都大于0了，则i、j也一定大于0，因此三者的和不可能等于0
		if nums[k] > 0 {
			break
		}

		// 3.2 指针k位置去重
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}

		// 3.3 开始固定指针i、j
		i := k + 1 // 指针i放在k之后的一位
		j := n - 1 // 指针j放在数组的最后
		// 当指针i、j不越界时
		for i < j {
			sum := nums[k] + nums[i] + nums[j]
			// 判断三者相加的和与0的关系
			if sum < 0 { // 3.3.1 三者和小于0
				// 指针i右移，寻找更大元素
				i++
				// 指针i去重
				for i < j && nums[i] == nums[i-1] {
					i++
				}
			} else if sum > 0 { // 3.3.2 三者和大于0
				// 指针j左移，寻找更小元素
				j--
				// 指针j去重
				for i < j && nums[j] == nums[j+1] {
					j--
				}
			} else { // 3.3.3 三者和等于0
				// 将当前结果加入结果集中
				oneOfRes := []int{nums[k], nums[i], nums[j]}
				res = append(res, oneOfRes)
				// 指针i、j分别移动，寻找下一个符合的组合
				i++
				j--
				// 指针i、j去重
				for i < j && nums[i] == nums[i-1] {
					i++
				}
				for i < j && nums[j] == nums[j+1] {
					j--
				}
			}
		}
	}

	// 4.返回结果集
	return res
}

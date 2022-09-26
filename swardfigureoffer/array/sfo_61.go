package array

func isStraight(nums []int) bool {
	// 最大值-最小值<5 且 除0外没有重复的元素

	// 哈希表
	var hashTable map[int]bool = make(map[int]bool)
	// 最大值最小值
	var minOfNums int = 14
	var maxOfNums int = -1
	// 遍历数组
	for _, num := range nums {
		// 如果是0，就跳过
		if num == 0 {
			continue
		}
		// 记录最大最小值，并判断是否重复
		if num > maxOfNums {
			maxOfNums = num
		}
		if num < minOfNums {
			minOfNums = num
		}
		if !hashTable[num] {
			hashTable[num] = true
		} else {
			return false
		}
	}
	// 跳出循环，说明没有重复的值，此时只需判断最大值和最小值的差值是否小于5即可
	if maxOfNums-minOfNums < 5 {
		return true
	} else {
		return false
	}
}

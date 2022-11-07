package array

func canJump(nums []int) bool {
	// 贪心算法

	/*
		1.数组中的元素表示在该位置可以跳跃的最大长度，因此从该位置可以跳到的最大下标=当前位置下标+当前位置的数组中元素
		2.从起点开始尝试着跳，跳跃长度从0开始逐次加1，直到尝试到可以跳到的最大下标为止
		3.每尝试跳一次，就看看跳到的新位置所能跳到的最大下标（新位置下标+新位置数组中元素）是否比老位置所能达到的最大下标大
			3.1 如果新位置比老位置所能达到的最大下标大，则更新所能达到的最大下标
			3.2 如果都跳出老位置能达到的最大下标了，还没找到能达到更大的最大下标的新位置，且老位置跳的最大下标无法达到终点，此时说明一定无法达到最后一个下标了，不再继续寻找，返回false
			3.3 如果最大下标超过了最后一个下标，说明能够到达最后一个下标，不再继续寻找，返回true如果最大下标超过了最后一个下标，说明能够到达最后一个下标，不再继续寻找，返回true
	*/

	// 1.初始化能到达的最大距离为0
	maxIndex := 0
	n := len(nums)

	// 2.从起点开始尝试着跳，跳跃长度从0开始逐次加1，直到尝试到可以跳到的最大下标为止
	for i := 0; i < n; i++ {
		// 判断此时是否跳出了能够达到的最大下标
		if i <= maxIndex { // 2.1 如果没有跳出最大下标
			// 如果新位置比老位置所能达到的最大下标大，则更新所能达到的最大下标
			maxIndex = maxInt(maxIndex, i+nums[i])
			// 如果最大下标超过了最后一个下标，说明能够到达最后一个下标，不再继续寻找，返回true
			if maxIndex >= n-1 {
				return true
			}
		} else { // 2.2 如果跳出了最大下标
			// 如果都跳出老位置能达到的最大下标了，还没找到能达到更大的最大下标的新位置
			// 且老位置跳的最大下标无法达到终点（没有返回true），此时说明一定无法达到最后一个下标了，不再继续寻找，返回false
			return false
		}
	}
	return false
}
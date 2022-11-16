package array

func maxProduct(nums []int) int {
	// 动态规划

	/*
		该题要求求出整个数组中乘积最大的非空连续子序列，而这个最大的连续子序列一定是以数组中的某个数为结尾的
		因此考虑使用动态规划，将求整个数组中乘积最大的非空连续子序列，拆解成求以数组中每个数为结尾的连续子序列的最大乘积，其中乘积最大的值就是我们要找的整个数组中乘积最大的连续子数组值
		1.子问题（状态）：但结尾数可能是正数也可能是负数，因此维护两个状态
			maxDp[i]为数组中以nums[i]为结尾的连续子序列的最大乘积
			minDp[i]为数组中以nums[i]为结尾的连续子序列的最小乘积
		2.起始状态：maxDp[0]和minDp[0]是以第0个元素为结尾的连续子序列乘积，此时最大乘积和最小乘积都是nums[0]
		3.状态转移：
			3.1 如果当前结尾nums[i]是正数，则：
				以nums[i-1]为结尾的最大连续子序列乘积(maxDp[i-1])乘正数后仍然是最大值，因此比较maxDp[i-1]*nums[i]和nums[i]的较大值，来选择nums[i]结尾的子集是否要乘上nums[i-1]结尾的子集还是nums[i]自己另起炉灶（单个数作为乘积子序列）会让以nums[i]为结尾的子序列乘积更大
				以nums[i-1]为结尾的最小连续子序列乘积(minDp[i-1])乘正数后仍然是最小值，因此比较minDp[i-1]*nums[i]和nums[i]的较小值，来选择nums[i]结尾的子集是否要乘上nums[i-1]结尾的子集还是nums[i]自己另起炉灶（单个数作为乘积子序列）会让以nums[i]为结尾的子序列乘积更小
			3.2 如果当前结尾nums[i]是负数，则：
				以nums[i-1]为结尾的最大连续子序列乘积(maxDp[i-1])乘负数后变成最小值，因此比较maxDp[i-1]*nums[i]和nums[i]的较小值，来选择nums[i]结尾的子集是否要乘上nums[i-1]结尾的子集还是nums[i]自己另起炉灶（单个数作为乘积子序列）会让以nums[i]为结尾的子序列乘积更小
				以nums[i-1]为结尾的最小连续子序列乘积(minDp[i-1])乘负数后变成最大值，因此比较minDp[i-1]*nums[i]和nums[i]的较大值，来选择nums[i]结尾的子集是否要乘上nums[i-1]结尾的子集还是nums[i]自己另起炉灶（单个数作为乘积子序列）会让以nums[i]为结尾的子序列乘积更大
			3.3 如果当前结尾nums[i]是0，则：
				任何数乘0都是0，因此无论是选择nums[i]结尾的子集乘上nums[i-1]结尾的子集还是nums[i]自己另起炉灶（单个数作为乘积子序列），以nums[i]为结尾的子序列乘积都是0，因此直接将maxDp[i]和minDp[i]置0
	*/

	// 1.特殊情况：
	n := len(nums)
	// 数组长度为0则直接返回0
	if n == 0 {
		return 0
	}
	// 数组长度为1则直接返回第一个数
	if n == 1 {
		return nums[0]
	}

	// 2.状态数组：maxDp[i]为数组中以nums[i]为结尾的连续子序列的最大乘积，minDp[i]为数组中以nums[i]为结尾的连续子序列的最小乘积
	maxDp := make([]int, n)
	minDp := make([]int, n)

	// 3.起始状态：maxDp[0]和minDp[0]是以第0个元素为结尾的连续子序列乘积，此时最大乘积和最小乘积都是nums[0]
	maxDp[0] = nums[0]
	minDp[0] = nums[0]

	// 4.状态转移：求以数组中每个数为结尾的连续子序列的最大乘积
	for i := 1; i < n; i++ {
		if nums[i] > 0 {
			// 4.1 如果nums[i]是正数，详见题解
			maxDp[i] = maxInt(maxDp[i-1]*nums[i], nums[i])
			minDp[i] = minInt(minDp[i-1]*nums[i], nums[i])
		} else if nums[i] < 0 {
			// 4.2 如果nums[i]是负数，详见题解
			maxDp[i] = maxInt(minDp[i-1]*nums[i], nums[i])
			minDp[i] = minInt(maxDp[i-1]*nums[i], nums[i])
		} else {
			// 4.3 如果nums[i]是0，详见题解
			maxDp[i] = 0
			minDp[i] = 0
		}
	}

	// 5.维持以nums[i]为结尾的连续子序列的最大乘积的状态数组maxDp中的最大值，就是我们要找的整个数组中乘积最大的连续子数组值
	res := maxOfIntArray(maxDp)

	// 6.返回结果
	return res
}

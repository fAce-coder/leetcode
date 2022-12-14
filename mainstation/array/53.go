package array

func maxSubArray(nums []int) int {
	// 动态规划

	/*
		该题要求整个数组中的最大连续子数组的和；而最终结果中的最大和子数组，肯定会以数组中某一个数为结尾
		考虑将整个数组中的最大和这一大问题，拆解成求以数组中每个元素为子数组结尾的最大子数组和 这些局部最优解，从这些子问题最优解中找到整个数组的最大和这一全局最优解
		因此考虑动态规划
		1.子问题(状态)：dp[i]为以该数组中索引i下的元素为结尾的连续子数组的最大和
		2.起始状态：dp[0]是数组中以第0个元素为结尾的连续子数组的最大和，此时该最大和一定是nums[0]
		3.状态转移：
			3.1 如果dp[i-1]的值是负数，则说明以nums[i-1]为结尾的子数组最大和是负数，则dp[i-1]+nums[i]的值<nums[i]，即：直接将nums[i]作为以nums[i]结尾的子数组最大和，会比nums[i]加上nums中前一个数为结尾的子数组最大和的情况大；因此此时选取nums[i]作为当前子问题的最优解
			3.2 如果dp[i-1]的值是正数，则说明以nums[i-1]为结尾的子数组最大和是正数，则dp[i-1]+nums[i]的值>nums[i]，即：将nums[i]加上nums中前一个数为结尾的子数组最大和，会比直接将nums[i]作为以nums[i]结尾的子数组最大和的情况大；因此此时选取dp[i-1]+nums[i]作为当前子问题的最优解
	*/

	// 1.初始化结果
	var res int

	// 2.特殊情况:如果数组的长度为1，则无需判断直接返回这一个数
	lengthOfNums := len(nums)
	if lengthOfNums == 1 {
		return nums[0]
	}

	// 3.定义与数组长度相同的状态转移数组
	// 子问题(状态)：dp[i]为以该数组中索引i下的元素为结尾的连续子数组的最大和
	dp := make([]int, lengthOfNums)

	// 4.起始状态：dp[0]是数组中以第0个元素为结尾的连续子数组的最大和，此时该最大和一定是nums[0]
	dp[0] = nums[0]

	// 5.状态转移
	for i := 1; i < lengthOfNums; i++ {
		// 判断以上一个数为结尾的子数组的和是否小于等于0
		if dp[i-1] <= 0 {
			// 5.1 如果dp[i-1]的值小于等于0，则说明以nums[i-1]为结尾的子数组最大和是负数，则dp[i-1]+nums[i]的值<nums[i]
			// 即：直接将nums[i]作为以nums[i]结尾的子数组最大和，会比nums[i]加上nums中前一个数为结尾的子数组最大和的情况大
			// 因此此时选取nums[i]作为当前子问题的最优解
			dp[i] = nums[i]
		} else {
			// 5.2 如果dp[i-1]的值是正数，则说明以nums[i-1]为结尾的子数组最大和是正数，则dp[i-1]+nums[i]的值>nums[i]
			// 即：将nums[i]加上nums中前一个数为结尾的子数组最大和，会比直接将nums[i]作为以nums[i]结尾的子数组最大和的情况大
			// 因此此时选取dp[i-1]+nums[i]作为当前子问题的最优解
			dp[i] = dp[i-1] + nums[i]
		}
	}

	// 6.最终结果（全局最优解）：状态数组中的最大值，即以该最大值索引所在的数结尾的子数组的和最大
	res = maxInt(dp...)

	// 7.返回结果
	return res
}

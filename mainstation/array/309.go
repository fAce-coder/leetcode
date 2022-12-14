package array

func maxProfit3(prices []int) int {
	// 动态规划

	/*
		本题要求整个周期中，在含冷冻期的情况下买卖多次股票所能得到的最大利润；即从周期第0天开始，到周期最后一天为止，所能得到的最大利润
		因此可以将整个周期中买卖多次股票这一大问题，拆分成求以第0天开始，以周期中每一天为止，买卖股票所能得到的最大利润这一子问题，因此考虑动态规划
		1.状态（子问题）：
			从周期第0天开始，到第i天为止所能得到的最大利润；又因为只有第i-1不持有股票且不是第i-1天卖出的（即不会导致第i天进入冷冻期），才可以在第i天买入股票，所以第i天可能有3种状态：当前不持有股票且处于冷冻期，当前不持有股票且不处于冷冻期，当前持有股票
			状态数组dp00[i]表示第i天不持有股票且不是第i天当天卖出的（不会导致下一天进入冷冻期），所能得到的最大利润
			状态数组dp01[i]表示第i天不持有股票且是第i天当天卖出的（会导致下一天进入冷冻期），所能得到的最大利润
			状态数组dp1[i]表示第i天持有股票时，所能得到的最大利润
		2.起始状态：
			如果第0天的状态是不持有股票，则说明第0天没有买入股票，或买入之后直接卖出了，则利润为0，此时dp00[0]=0,dp01[0]=0
			如果第0天的状态是持有股票，则说明第0天买入了股票，则利润为负的买入股票的钱，此时dp1[0]=-prices[0]
		3.状态转移：
			当第i天不持有股票，且股票不是第i天当天卖出的：
				说明是第i-1天已经不持有股票了，第i天也没有买入，则第i天为止的利润=第i-1天为止的利润，而由于第i-1天也不持有股票，因此第i-1天的最大利润取两种状态（是否是第i-1天当天卖出的）的较大值，即：dp00[i]=max(dp00[i-1],dp01[i-1])
			当第i天不持有股票，且股票是第i天当天卖出的：
				说明是第i-1天还持有股票，是第i天卖出的，则第i天为止的利润=第i-1天为止的利润+第i天卖出股票赚的钱，即：dp01[i]=dp1[i-1]+prices[i]
			当第i天持有股票：
				如果是第i-1天已经持有股票了，第i天也没有卖出，则第i天为止的利润=第i-1天为止的利润，即dp1[i]=dp1[i-1]
				如果是第i-1天不持有股票，是在第i天买入的，因此股票不能是第i-1天当天卖的（第i天不处于冷冻期），则第i天为止的利润=第i-1天为止的利润-第i天买入股票花的钱，即：dp1[i]=dp00[i-1]-prices[i]
				因此第i天持有股票的最大利润是上述两种情况的较大值
			因为若想得到整个周期中买卖股票的最大利润，则最后一天一定不能持有股票，因此选取最后一天dp00和dp01的较大值，就是我们要求的最终结果
	*/

	// 1.状态：
	// 从周期第0天开始，到第i天为止所能得到的最大利润；又因为只有第i-1不持有股票且不是第i-1天卖出的（即不会导致第i天进入冷冻期），才可以在第i天买入股票，所以第i天可能有3种状态：当前不持有股票且处于冷冻期，当前不持有股票且不处于冷冻期，当前持有股票
	n := len(prices)
	dp00 := make([]int, n) // 状态数组dp00[i]表示第i天不持有股票且不是第i天当天卖出的（不会导致下一天进入冷冻期），所能得到的最大利润
	dp01 := make([]int, n) // 状态数组dp01[i]表示第i天不持有股票且是第i天当天卖出的（会导致下一天进入冷冻期），所能得到的最大利润
	dp1 := make([]int, n)  // 状态数组dp1[i]表示第i天持有股票时，所能得到的最大利润

	// 2.起始状态
	dp00[0] = 0 // 如果第0天的状态是不持有股票，则说明第0天没有买入股票，或买入之后直接卖出了，则利润为0，此时dp00[0]=0,dp01[0]=0
	dp01[0] = 0
	dp1[0] = -prices[0] // 如果第0天的状态是持有股票，则说明第0天买入了股票，则利润为负的买入股票的钱，此时dp1[0]=-prices[0]

	// 3.状态转移
	for i := 1; i < n; i++ {
		// 3.1 当第i天不持有股票，且股票不是第i天当天卖出的：
		// 说明是第i-1天已经不持有股票了，第i天也没有买入，则第i天为止的利润=第i-1天为止的利润，而由于第i-1天也不持有股票，因此第i-1天的最大利润取两种状态（是否是第i-1天当天卖出的）的较大值，即：dp00[i]=max(dp00[i-1],dp01[i-1])
		dp00[i] = maxInt(dp00[i-1], dp01[i-1])
		// 3.2 当第i天不持有股票，且股票是第i天当天卖出的：
		// 说明是第i-1天还持有股票，是第i天卖出的，则第i天为止的利润=第i-1天为止的利润+第i天卖出股票赚的钱，即：dp01[i]=dp1[i-1]+prices[i]
		dp01[i] = dp1[i-1] + prices[i]
		// 3.3 当第i天持有股票：
		// 如果是第i-1天已经持有股票了，第i天也没有卖出，则第i天为止的利润=第i-1天为止的利润，即dp1[i]=dp1[i-1]
		// 如果是第i-1天不持有股票，是在第i天买入的，因此股票不能是第i-1天当天卖的（第i天不处于冷冻期），则第i天为止的利润=第i-1天为止的利润-第i天买入股票花的钱，即：dp1[i]=dp00[i-1]-prices[i]
		// 因此第i天持有股票的最大利润是上述两种情况的较大值
		dp1[i] = maxInt(dp1[i-1], dp00[i-1]-prices[i])
	}

	// 4.因为若想得到整个周期中买卖股票的最大利润，则最后一天一定不能持有股票，因此选取最后一天dp00和dp01的较大值，就是我们要求的最终结果
	return maxInt(dp00[n-1], dp01[n-1])
}

package array

func nextPermutation(nums []int) {
	/*
		本题要得到下一个更大，但变大的幅度最小的排列
		1.因此，首先考虑第一个条件，如何将排列变大：
			将当前排列后面的较大的数j（后续称作大数），与前面的较小的数i（后续称作小数）进行交换，即可使得排列变大
		2.接下来考虑第二个条件，如何让变大的幅度最小：
			因为大数和小数交换的位置越靠后，即越靠近个位数，其增大的幅度越小；因此要使得小数尽量靠右，而大数尽可能的小
			2.1 先找小数：
				因为既要满足小数尽量靠右，又要满足小数<大数，且小数在大数的前边，因此，从前向后找到最后一个升序的组合i,i+1，使得nums[i]<nums[i+1]（为了提高性能，可以从后向前找）
				此时索引i位置就是要交换的小数（因为从i到i+1是最后一个升序，所以从i+1到end都是降序，不再存在满足上3个条件的大数和小数）
			2.2 再找大数：
				小数的交换位置已经确定，因此为了使得变大的幅度最小，所以大数越小越好
				因此从i+1到end索引中，找到比小数大，但大的幅度最小的数，则那个数就是待交换的大数（因为i+1到end是降序，因此可以从end开始倒着找）
			2.3 交换后，将i+1到end变成升序
				因为大数和小数交换后，排列已经变大，但要求其变大的最小，则需要将i+1到end升序，使小数i索引后面的数是最小状态
				因为交换前大数是比小数大一点的数，因此交换后，i+1到end还是降序排列，所以要将其升序，只需要将i+1到end倒置即可
	*/

	// 1.特殊情况，如果数组中只有1个或0个数，则下一个排列固定不变
	n := len(nums)
	if n <= 1 {
		return
	}

	// 2.先找小数
	// 从后向前找到最后一个升序的组合i,i+1，使得nums[i]<nums[i+1]
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	// 3.如果小数没越界，再找大数；如果小数越界了，说明整个数组都是降序，说明此时是最大状态，将整个数组改成升序即可
	// 小数的交换位置已经确定，因此为了使得变大的幅度最小，所以大数越小越好
	if i >= 0 {
		// 3.1 从i+1到end索引中，找到比小数大，但大的幅度最小的数，则那个数就是待交换的大数（因为i+1到end是降序，因此可以从end开始倒着找）
		j := n - 1
		for j > i && nums[j] <= nums[i] {
			j--
		}
		// 3.2 交换大数和小数
		nums[i], nums[j] = nums[j], nums[i]
	}

	// 4.将i+1到end位置的数组变成升序
	// 如果小数没越界：
	// 		因为大数和小数交换后，排列已经变大，但要求其变大的最小，则需要将i+1到end升序，使小数i索引后面的数是最小状态
	//		因为交换前大数是比小数大一点的数，因此交换后，i+1到end还是降序排列，所以要将其升序，只需要将i+1到end倒置即可
	// 如果小数越界了：
	// 		说明整个数组都是降序，说明此时是最大状态，将整个数组改成升序即可
	// 		因为是降序，同样只需将i+1到end倒置即可
	left := i + 1
	right := n - 1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
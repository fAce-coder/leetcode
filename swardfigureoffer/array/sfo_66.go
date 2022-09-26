package array

func constructArr(a []int) []int {
	// 辅助数组

	/*
	   每个数都是左边的乘积，再乘上右边的乘积
	   所以创建两个数组，一个存左边的乘积，一个存右边的乘积
	   最后将两个数组的对应位置相乘，就是结果数组
	*/

	// 特殊情况
	var length int = len(a)
	if length == 0 {
		return []int{}
	}
	// 创建左右乘积数组
	var leftArr []int = make([]int, length)
	var rightArr []int = make([]int, length)
	var res []int = make([]int, length)
	// 构建左乘积数组
	leftArr[0] = 1
	for i := 1; i < length; i++ {
		leftArr[i] = leftArr[i-1] * a[i-1]
	}
	// 构建右乘积数组
	rightArr[length-1] = 1
	for i := length - 2; i > -1; i-- {
		rightArr[i] = rightArr[i+1] * a[i+1]
	}
	// 左右乘积数组相乘，得出结果
	for i := 0; i < length; i++ {
		res[i] = leftArr[i] * rightArr[i]
	}
	// 将结果返回
	return res
}

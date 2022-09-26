package array

func singleNumber1(nums []int) int {
	// 位运算

	/*
	   1个数字出现1次，其他数字都出现3次
	   说明如果将所有数的二进制位相加，则1的个数只能是两种情况：模3余1，模3余0
	   将所有数的二进制位对位相加，并存储在32位数组中，然后模3，就能得出那个出现1次的数的二进制形式
	   将这个二进制位还原成原始数字并输出
	*/

	// 定义储存32位二进制数的数组
	var bitNums []int = make([]int, 32)
	// 遍历数组，将所有数转换成二进制形式，并加在数组的对应位置
	for _, num := range nums {
		for i := 31; i > -1; i-- {
			bitNums[i] += num & 1
			num >>= 1
		}
	}
	// 遍历bit数组，对每个数模3，获取那个只出现一次的数的二进制数组
	for i, _ := range bitNums {
		bitNums[i] %= 3
	}
	// 遍历bit数组，还原这个数
	var res int = 0
	for _, num := range bitNums {
		res <<= 1
		res |= num
	}
	// 返回结果
	return res
}

func singleNumber2(nums []int) int {
	// HashMap

	/*
	   利用HashMap去除重复元素
	   以数字为键，出现次数为值
	   遍历原数组，将遍历到的数字在HashMap中出现次数加1
	   遍历HashMap，返回值为1的键
	*/

	// 建立HashMap
	var HashMap map[int]int = make(map[int]int)
	// 遍历原数组，将遍历到的数字在HashMap中出现次数加1
	for _, num := range nums {
		HashMap[num] += 1
	}
	// 遍历HashMap，输出值为1的键
	var res int
	for key, value := range HashMap {
		if value == 1 {
			res = key
			break
		}
	}
	// 返回结果
	return res
}

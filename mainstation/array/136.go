package array

func singleNumber1(nums []int) int {
	// 位运算

	/*
		该题要求找出数组中唯一一个出现了2次的数字，这种有去重相关思想的题，首先就能想到最简单的哈希表，在golang中即map
		但该题额外要求是不使用额外空间，因此考虑结合位运算中的异或运算的特性
		异或特性：两个相同的数字异或的结果是0，任意数字与0的异或结果都是数字本身，而异或操作又满足交换律，即a^(b^c)=(a^b)^c
		因此，只要将数组中所有的数字都进行异或操作，则数组中出现2次的数字会结合变成0，最后那个出现1次的数字与这些0异或会得到自身，这样就可以得到这个出现了一次的数字
	*/

	// 1.初始化结果是0(与任何数字异或都不影响那个值)
	res := 0

	// 2.遍历数组，将数组中的所有值进行异或操作，得到结果
	for _, num := range nums {
		res ^= num
	}

	// 3.将结果返回
	return res
}

func singleNumber2(nums []int) int {
	// 哈希表

	/*
		遍历数组，每遇到一个数字，就将该数字存入map的键，将其值+1，最后遍历map，找出那个值是1的键即可
	*/

	// 1.初始化结果
	var res int

	// 2.初始化哈希表
	hashMap := make(map[int]int)

	// 3.每遇到一个数字，就将该数字存入map的键，将其值+1，表示此时出现了一次
	for _, num := range nums {
		hashMap[num]++
	}

	// 4.最后遍历map，找出那个值是1的键即可
	for key, value := range hashMap {
		if value == 1 {
			res = key
			break
		}
	}

	// 5.将结果返回
	return res
}

package array

func singleNumbers(nums []int) []int {
	// 位运算

	/*
	   两个相同的数字，异或结果是0
	   在数组中，除了2个数字外，其他都出现了2次，因此所有数字异或的结果=这2个数字异或的结果
	   将数组异或结果，从低位到高位，找出第一个为1的位（说明该位二者不相同，一个为0，一个为1）
	   要找出这两个数，只需要按照这个不相同的位，将数组中的数分成两组，再分组进行异或，此时即可抵消，找到结果
	*/

	// 将数组中的数字全部异或，得出这两个数字的异或结果
	var temp1 int = 0
	for _, num := range nums {
		temp1 = temp1 ^ num
	}
	// 从低到高，找到第一个为1的位
	var temp2 int = 1
	for temp1&temp2 == 0 {
		temp2 = temp2 << 1
	}
	// 将数组中的数分组异或，得出结果
	var res1 int = 0
	var res2 int = 0
	for _, num := range nums {
		if num&temp2 == 0 {
			res1 = res1 ^ num
		} else {
			res2 = res2 ^ num
		}
	}
	// 将两个结果返回
	return []int{res1, res2}
}

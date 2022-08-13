package array

func findRepeatNumber(nums []int) int {
	// 哈希表，map版，效率更高

	// 创建一个hashMap
	var hashMap map[int]bool = make(map[int]bool)

	for _, num := range nums {
		// 如果数字在map的key里，就返回这个数
		if hashMap[num] {
			return num
		}
		// 如果数字不在map的key里，就把这个数加入map的key
		hashMap[num] = true
	}
	// 如果到最后也没找到重复的数，则返回-1
	return -1
}

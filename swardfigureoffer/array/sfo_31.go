package array

func validateStackSequences(pushed []int, popped []int) bool {
	// 辅助栈

	// 创建辅助栈
	// var length int = len(pushed)
	var temp []int = make([]int, 0)
	// pop栈索引
	var i int = 0
	// 按照push栈入栈
	for _, num := range pushed {
		// 当辅助栈栈顶元素与pop栈的首元素相同，且辅助栈不为空，辅助栈出栈，pop栈右移
		temp = append(temp, num)
		for len(temp) != 0 && temp[len(temp)-1] == popped[i] {
			temp = temp[:len(temp)-1]
			i++
		}
	}
	// 全部入栈完成后，如果辅助栈为空，说明序列相同，否则不同
	if len(temp) == 0 {
		return true
	} else {
		return false
	}
}

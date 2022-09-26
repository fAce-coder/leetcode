package array

func findContinuousSequence(target int) [][]int {
	// 滑动窗口

	// 定义返回结果
	var res [][]int = make([][]int, 0)
	// 特殊情况：连续正整数序列，且至少含有2个数 -> 1 + 2 = 3,因此target至少>=3
	if target < 3 {
		return res
	}
	// 快慢双指针
	var slow int = 1
	var fast int = 2
	var sumOfSlowAndFast int = slow + fast
	// 滑动窗口开始移动
	for slow < fast {
		// sum < target，快指针右移，增大滑动窗口
		if sumOfSlowAndFast < target {
			fast++
			sumOfSlowAndFast += fast
			// sum > target,慢指针右移，减小滑动窗口
		} else if sumOfSlowAndFast > target {
			sumOfSlowAndFast -= slow
			slow++
			// sum = target,将结果加入，快指针右移，增大滑动窗口
		} else {
			// 临时数组
			temp := make([]int, 0)
			for i := slow; i <= fast; i++ {
				temp = append(temp, i)
			}
			// 加入结果中
			res = append(res, temp)
			// 右指针右移
			fast++
			sumOfSlowAndFast += fast
		}
	}
	// 将结果返回
	return res
}

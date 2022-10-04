package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// maxInt
//
//	@Description: 求两个整型数据的最大值
func maxInt(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

// absInt
//
//	@Description: 求一个数的绝对值
func absInt(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}

// reverseIntSlice
//
//	@Description: 翻转切片
func reverseIntSlice(nums []int) {
	i := 0
	j := len(nums) - 1

	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

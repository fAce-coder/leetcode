package array

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	// 递归+剪枝

	/*
		这类问题先画出树状图，以target为根节点，创建分支时根节点与候选数组中的每一个数（树的边）做减法来得到一个子节点（剩余目标值）
		再以子节点为根节点依次类推，尝试将数减到0（减到0或负数时停止）
		最终，所有从根节点到节点0的路径就是其中的一个结果
		为了避免重复结果和额外查找，将候选数组排序并进行剪枝操作
	*/

	// 1.特殊情况：候选数组为0，直接返回
	if len(candidates) == 0 {
		return [][]int{}
	}

	// 2.初始化结果集
	res := make([][]int, 0)

	// 3.对候选数组进行从小到大排序，方便剪枝
	sort.Sort(sortableNums(candidates))

	// 4.定义递归函数，来从根节点开始寻找目标
	var dfs func(int, []int, int)
	// begin：在每一层中树的边，即用来与上级节点相减的候选数组中的某个数
	// path：从最顶层的根节点开始的每一条符合条件（终点是0）的路径
	// target：每一层中树的节点，即上级节点与候选数组中的某个数相减后的结果，即剩余目标值
	dfs = func(begin int, path []int, target int) {
		// 4.1 终止条件：如果剩余目标值等于0，则说明刚好组合完，是满足题目要求的一组数
		if target == 0 {
			// 将这组数加入结果集，并将树中的该分支退出
			res = append(res, path)
			return
		}
		// 4.2 让根节点，分别与候选数组中的每一个数相减，来得到子节点，即创建多个分支来递归查找满足条件的分支
		for i := begin; i < len(candidates); i++ {
			// 4.2.1 求出子节点（即剩余目标值）
			residue := target - candidates[i]
			// 4.2.2 如果剩余目标值小于0，说明该分支不符合
			// 且因为数组是从小到大排序的，所以如果继续向后找要相减的数，只会导致剩余目标值都小于0且会更小，因此跳出循环，不再向后查找（剪枝）
			if residue < 0 {
				break
			}
			// 4.2.3 将当前边（用来与上级节点相减的候选数组中的某个数）加入这条路径中
			// 为防止切片扩容相互影响，使分支相互独立，因此新创建一个切片，将原切片追加后再追加新值
			newPath := make([]int, 0)
			newPath = append(newPath, path...)
			newPath = append(newPath, candidates[i])
			// 4.2.4 递归的寻找当前分支的下一个边
			dfs(i, newPath, residue)
		}
	}

	// 5.执行递归函数，从候选数组的第一个数作为第一个边；起始分支路径为空；起始目标值为完整目标值
	dfs(0, []int{}, target)

	// 6.返回结果
	return res
}

package array

// 参考答案：https://leetcode.cn/problems/subsets/solution/c-zong-jie-liao-hui-su-wen-ti-lei-xing-dai-ni-gao-/

func subsets(nums []int) [][]int {
	// 递归+回溯

	/*
		这类递归回溯的问题，首先画出一个树状图，以空结果集为根节点，创建分支时根节点将候选数组中的每一个数（树的边）加入到临时结果集中，来得到一个子节点（临时结果集，即nums的其中一个子集）
		再以子节点为根节点依次类推，因为要求子集不能重复，因此在创建分支时，只找候选数组中当前索引后面的数作为当前分支下一个加入到临时结果集中的数（树的边）
		最终，树中的每一个分支的每一个节点就是一个临时结果集（nums数组的其中一个子集）
	*/

	// 1.初始化最终结果集和临时结果集
	res := make([][]int, 0)
	path := make([]int, 0)

	// 2.递归回溯函数
	var dfs func(int, []int)
	// begin：用来标记当前分支的当前节点（临时结果集）已经找到了哪个索引，为避免重复子集，因此从当前节点再创建分支时，就只找nums数组中begin索引后面的数
	// path：记录树中的每一个临时结果集（nums中的一个子集，树的节点）
	dfs = func(begin int, path []int) {
		// 2.1 将树中的每一个节点（临时结果集，nums的一个子集）加入到最终结果集中
		newPath := make([]int, 0)
		newPath = append(newPath, path...)
		res = append(res, newPath)
		// 2.2 以当前节点为根节点，将候选数组中的每一个数（树的边）加入到临时结果集中并创建一个分支，来得到一个子节点
		for i := begin; i < len(nums); i++ {
			// 2.2.1 将其中一个数作为边创建一个分支，来得到一个子节点
			path = append(path, nums[i])
			// 2.2.2 在当前分支中递归寻找其子节点的子节点，但只寻找比当前索引大的数
			dfs(i+1, path)
			// 2.2.3 回溯，以便在下一轮循环中将下一个数作为树的边加入到临时结果集中以创建一个新的分支
			path = path[:len(path)-1]
		}
	}

	// 3.执行递归，初始化索引为0，path为空
	dfs(0, path)

	// 4.返回结果
	return res
}

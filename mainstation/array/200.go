package array

func numIslands(grid [][]byte) int {
	// 递归

	/*
		本题要找到岛屿数量，即通过上下左右相连接的数字1的块数
		因此我们可以遍历这个二维网格中的每一个元素，遇到1就说明当前是陆地，陆地数量加1，同时为了防止重复查找，将当前1所在的一整块陆地全都变成水（数字0）
		因为在消除当前一整块陆地时，是从当前点逐渐向上下左右蔓延直到不是陆地（数字0）为止，因此考虑使用递归
		1.终止条件：
			在消除陆地时，只要越界或者找到了不是陆地（数字0），说明这整块待消除陆地在当前方向上已经没有了，就终止查找并返回
		2.本次递归中要做的事：
			将当前这块陆地（数字1）变成水（数字0），来防止重复查找
		3.递归：
			继续朝当前位置的上下左右进行递归，查找相连的陆地（数字1），将其置为水（数字0）
	*/

	// 1.初始化变量
	// 二维网格的长m和宽n
	m := len(grid)
	n := len(grid[0])
	// 初始化岛屿数量为0
	res := 0

	// 2.递归函数：用来消除一整块的陆地
	var dfs func(int, int)
	// i：当前所在的行索引
	// j：当前所在的列索引
	dfs = func(i int, j int) {
		// 2.1 终止条件：在消除陆地时，只要越界或者找到了不是陆地，说明这整块待消除陆地在当前方向上已经没有了，就终止查找并返回
		if (i < 0 || i >= m) || (j < 0 || j >= n) || grid[i][j] != '1' {
			return
		}
		// 2.2 本次递归中要做的事：将当前这块陆地（数字1）变成水（数字0），来防止重复查找
		grid[i][j] = '0'
		// 2.3 递归：继续朝当前位置的上下左右进行递归，查找相连的陆地（数字1），将其置为水（数字0）
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}

	// 3.遍历这个二维网格中的每一个元素
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				// 遇到1就说明当前是陆地，陆地数量加1
				res++
				// 为了防止重复查找，将当前1所在的一整块陆地全都变成水（数字0）
				dfs(i, j)
			}
		}
	}

	// 4.返回结果
	return res
}

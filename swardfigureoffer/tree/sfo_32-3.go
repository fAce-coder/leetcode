package tree

func levelOrder3(root *TreeNode) [][]int {
	// 迭代

	/*
		从上到下打印二叉树，即层序遍历二叉树
		由于同一层的节点要按照从左到右的顺序打印，因此考虑用辅助队列来实现
		将根节点入队，并依次出队，每次出队后，就将其孩子按照左右孩子的顺序入队，直到所有节点都出队完毕
		因为每一层打印到一行，所以要控制每次只出队一行，新增标志位，记录每一行的长度
		因为每打印一层，就要交换打印顺序，因此新增标志位，记录打印顺序
	*/

	// 1.特殊情况：根节点是空，则直接返回空切片
	if root == nil {
		return [][]int{}
	}

	// 2.初始化结果集和队列
	res := make([][]int, 0)
	queue := make([]*TreeNode, 0)

	// 3.根节点入队
	queue = append(queue, root)

	// 4.判断打印顺序的标志位，初始化为1，即：从左向右打印
	order := 1

	// 5.当队列不是空，循环出队入队
	for len(queue) != 0 {
		// 5.1 记录当前层的节点个数
		lengthOfCurrentLayer := len(queue)
		// 5.2 记录当前层的结果
		resOfCurrentLayer := make([]int, 0)
		// 5.3 以当前层的节点个数为出队依据，进行循环出队
		for i := 0; i < lengthOfCurrentLayer; i++ {
			// 5.3.1 出队
			node := queue[0]
			queue = queue[1:]
			// 5.3.2  将出队节点加入当前层结果集
			resOfCurrentLayer = append(resOfCurrentLayer, node.Val)
			// 5.3.3 如果出队节点有左右孩子，就按照先左孩子后右孩子的顺序将节点入队
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// 5.4 判断打印顺序标志位
		if order == 1 {
			// 5.4.1 如果是1，则从左向右打印该层
			res = append(res, resOfCurrentLayer)
		} else {
			// 5.4.2 如果不是1，则从右向左打印该层，即翻转该层结果再加入最终结果集
			reverseIntSlice(resOfCurrentLayer)
			res = append(res, resOfCurrentLayer)
		}
		// 5.5 将order标志位置为负，表示下一层反向打印
		order = -order
	}

	// 6.返回最终结果集
	return res
}

func reverseIntSlice(nums []int) {
	// 反转切片
	i := 0
	j := len(nums) - 1

	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

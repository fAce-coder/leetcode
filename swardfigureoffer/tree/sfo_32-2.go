package tree

func levelOrder2(root *TreeNode) [][]int {
	// 迭代

	/*
		从上到下打印二叉树，即层序遍历二叉树
		由于同一层的节点要按照从左到右的顺序打印，因此考虑用辅助队列来实现
		将根节点入队，并依次出队，每次出队后，就将其孩子按照左右孩子的顺序入队，直到所有节点都出队完毕
		因为每一层打印到一行，所以要控制每次只出队一行，新增标志位，记录每一行的长度
	*/

	// 1.特殊情况：根节点是空，直接返回空切片
	if root == nil {
		return [][]int{}
	}

	// 2.初始化结果集和队列
	res := make([][]int, 0)
	queue := make([]*TreeNode, 0)

	// 3.将根节点加入队列
	queue = append(queue, root)

	// 4.当队列不为空，循环出队入队
	for len(queue) != 0 {
		// 4.1 记录当前层的节点个数
		lengthOfCurrentLayer := len(queue)
		// 4.2 记录当前层的结果
		resOfCurrentLayer := make([]int, 0)
		// 4.3 以当前层的节点个数为出队依据，进行循环出队
		for i := 0; i < lengthOfCurrentLayer; i++ {
			// 4.3.1 出队
			node := queue[0]
			queue = queue[1:]
			// 4.3.2 将出队节点加入当前层结果集
			resOfCurrentLayer = append(resOfCurrentLayer, node.Val)
			// 4.3.3 如果出队节点有左右孩子，就按照先左孩子后右孩子的顺序将节点入队
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// 4.4 将当前层的结果集加入到最终结果集中
		res = append(res, resOfCurrentLayer)
	}

	// 5.返回最终结果集
	return res
}

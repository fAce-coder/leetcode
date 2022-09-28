package tree

func levelOrder1(root *TreeNode) []int {
	// 辅助队列

	/*
		从上到下打印二叉树，即层序遍历二叉树
		由于同一层的节点要按照从左到右的顺序打印，因此考虑用辅助队列来实现
		将根节点入队，并依次出队，每次出队后，就将其孩子按照左右孩子的顺序入队，直到所有节点都出队完毕
	*/

	// 1.特殊情况：如果根节点是空，就返回一个空切片
	if root == nil {
		return []int{}
	}

	// 2.初始化结果集和队列
	res := make([]int, 0)
	queue := make([]*TreeNode, 0)

	// 3.将根节点加入队列
	queue = append(queue, root)

	// 4.当队列不为空时，循环出队入队
	for len(queue) != 0 {
		// 4.1 出队
		node := queue[0]
		queue = queue[1:]
		// 4.2 将出队节点的值加入到结果集中
		res = append(res, node.Val)
		// 4.3 如果出队节点存在左右孩子，就按照左右孩子的顺序入队
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	// 5.返回结果集
	return res
}

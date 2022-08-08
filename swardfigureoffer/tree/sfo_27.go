package tree

func mirrorTree1(root *TreeNode) *TreeNode {
	// 迭代

	/*
		通过队列实现迭代
		先将头结点入队，然后逐层进行翻转
	*/

	// 1.特殊情况，当二叉树的节点数量是0，则直接返回
	if root == nil {
		return nil
	}
	// 2.新建一个存放二叉树节点的队列，将根节点加进去
	var treeNodesQueue []*TreeNode = make([]*TreeNode, 0)
	treeNodesQueue = append(treeNodesQueue, root)
	// 3.当队列不为空，循环入队，将左右子树进行翻转
	for len(treeNodesQueue) != 0 {
		// 3.1 出队
		node := treeNodesQueue[0]
		treeNodesQueue = treeNodesQueue[1:]
		// 3.2 左右子树互换
		node.Left, node.Right = node.Right, node.Left
		// 3.3 左右子树入队，以便后续将他们的左右子树互换
		if node.Left != nil {
			treeNodesQueue = append(treeNodesQueue, node.Left)
		}
		if node.Right != nil {
			treeNodesQueue = append(treeNodesQueue, node.Right)
		}
	}
	// 4.将翻转之后的树的根节点返回
	return root

}

func mirrorTree2(root *TreeNode) *TreeNode {
	// 递归

	/*
		该函数的功能就是翻转树的左右子树
		所以在翻转完头结点之后，可以递归的使用该函数继续对其左右子树进行翻转
		终止条件：node是nil的时候就返回，说明根节点是空，或者遍历到树的叶子结点是空了
		本次递归内的操作：将传入节点的左右子树进行翻转
		递归：对其左右子树分别执行递归函数
	*/

	// 1.终止条件：node是nil的时候就返回，说明根节点是空，或者遍历到树的叶子结点是空了
	if root == nil {
		return nil
	}
	// 2.本次递归内的操作：将传入节点的左右子树进行翻转
	root.Left, root.Right = root.Right, root.Left
	// 3.递归：对其左右子树分别执行递归函数
	_ = mirrorTree2(root.Left)
	_ = mirrorTree2(root.Right)
	// 4.将根节点返回
	return root
}

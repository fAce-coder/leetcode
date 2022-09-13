package tree

func maxDepth1(root *TreeNode) int {
	// 递归

	/*
		终止条件：root为nil，说明根节点为空，或者找到最后的空叶子节点，返回深度0
		在本次递归中要做的事：当前节点不为空的时候，就将深度加1
		递归：递归的判断左右子树的深度，选出深度最大的，就是树的深度
	*/

	// 1.终止条件：root为nil，说明根节点为空，或者找到最后的空叶子节点，返回深度0
	if root == nil {
		return 0
	}
	// 2.在本次递归中要做的事：当前节点不为空的时候，就将深度加1
	// 3.递归：递归的判断左右子树的深度，选出深度最大的，就是树的深度
	return 1 + maxInt(maxDepth1(root.Left), maxDepth1(root.Right))
}

func maxDepth2(root *TreeNode) int {
	// 迭代

	/*
		利用队列实现迭代
		每次入队一层
		将一层出队，将深度加1，并将每个出队节点的左右孩子加入（如果有）
	*/

	// 1.特殊情况，当根节点为空时，此时树的深度为0
	if root == nil {
		return 0
	}

	// 2.初始化
	// 2.1 初始化树的深度为0
	var depthOfTree int = 0
	// 2.2 初始化空队列
	var treeNodeQueue []*TreeNode = make([]*TreeNode, 0)
	// 2.3 将树的根节点加入队列
	treeNodeQueue = append(treeNodeQueue, root)

	// 3.当队列不为空，循环出队
	for len(treeNodeQueue) != 0 {
		// 3.1 获取当前队列的长度
		// 这是保证每次只出队一层的关键
		// 如果不指定该变量，后续会不停向队列中加入节点，无法控制只出队一层
		curLengthOfQueue := len(treeNodeQueue)
		// 3.2 对一层的节点进行出队入队
		for i := curLengthOfQueue; i > 0; i-- {
			// 3.2.1 出队
			node := treeNodeQueue[0]
			treeNodeQueue = treeNodeQueue[1:]
			// 3.2.2 将其左右孩子加入（如果存在）
			if node.Left != nil {
				treeNodeQueue = append(treeNodeQueue, node.Left)
			}
			if node.Right != nil {
				treeNodeQueue = append(treeNodeQueue, node.Right)
			}
		}
		// 3.3 每出队一层，树的深度加1
		depthOfTree++
	}

	// 4.返回树的深度
	return depthOfTree
}

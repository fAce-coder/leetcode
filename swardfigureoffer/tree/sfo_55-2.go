package tree

func isBalanced(root *TreeNode) bool {
	// 	递归

	/*
		平衡二叉树的定义：(必须同时满足以下两个条件的二叉树才是平衡二叉树)
			左子树和右子树的深度之差不能超过1（<=1），此处可以引用sfo_55-1中的二叉树深度函数
			左子树和右子树也是平衡二叉树
		1.终止条件：
			如果根节点不存在，则返回True；
			如果到某个分支的左右子树深度之差大于1，或者其左右子树不是平衡二叉树，则说明不平衡，返回False；反之返回True
		2.本次递归要做的事：判断以当前节点为根节点的树，是否是平衡二叉树(左右子树深度之差是否<=1)
		3.递归调用：递归判断当前节点的左右子树是否是平衡二叉树
	*/

	// 1.终止条件：如果根节点不存在，则返回True
	if root == nil {
		return true
	}

	// 2.本次递归要做的事：判断以当前节点为根节点的树，是否是平衡二叉树(左右子树深度之差是否<=1)
	// 3.递归调用：递归判断当前节点的左右子树是否是平衡二叉树
	if absInt(maxDepth1(root.Left)-maxDepth1(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right) {
		return true
	} else {
		return false
	}
}

package tree

func lowestCommonAncestor1(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	// 	迭代

	/*
		二叉搜索树特性：某个节点，比该节点小的都在该节点的左子树，比该节点大的都在节点的右子树
		因此，当某个节点在p，q之间时，该节点就是p，q的最近公共祖先
		1.如果p，q都<当前节点，说明p，q都在当前节点的左子树，则当前节点移动至其左孩子
		2.如果p，q都>当前节点，说明p，q都在当前节点的右子树，则当前节点移动至其右孩子
		3.如果当前节点夹在p，q之间，则当前节点是p，q的最近公共祖先，可能有以下三种情况：
			p，q在最近公共祖先两侧：p>当前节点且q<当前节点，或者p<当前节点且q>当前节点
			p是公共祖先：p=当前节点
			q是公共祖先：q=当前节点
	*/

	// 当root不为空
	for root != nil {
		if p.Val < root.Val && q.Val < root.Val {
			// 1.如果p，q都<当前节点，说明p，q都在当前节点的左子树，则当前节点移动至其左孩子
			root = root.Left
		} else if p.Val > root.Val && q.Val > root.Val {
			// 2.如果p，q都>当前节点，说明p，q都在当前节点的右子树，则当前节点移动至其右孩子
			root = root.Right
		} else {
			// 3.如果当前节点夹在p，q之间，则当前节点是p，q的最近公共祖先
			return root
		}
	}
	// 特殊条件：root为空，返回nil
	return nil
}

func lowestCommonAncestor2(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	// 	递归

	/*
		1.终止条件：当根节点是空的时候返回nil；如果p，q不是同时大于或小于某个节点，则说明当前节点是最近公共祖先，将其返回
		2.本次递归中做的事：找到当前节点为根节点的树中的p，q的最近公共祖先
		3.递归：当p，q都小于或大于当前节点时，分别向其左右子树递归查找
	*/

	if root == nil {
		return nil
	}

	if p.Val < root.Val && q.Val < root.Val {
		// 1.如果p，q都<当前节点，说明p，q都在当前节点的左子树，则当前节点移动至其左孩子
		return lowestCommonAncestor2(root.Left, p, q)
	} else if p.Val > root.Val && q.Val > root.Val {
		// 2.如果p，q都>当前节点，说明p，q都在当前节点的右子树，则当前节点移动至其右孩子
		return lowestCommonAncestor2(root.Right, p, q)
	} else {
		// 3.如果当前节点夹在p，q之间，则当前节点是p，q的最近公共祖先
		return root
	}
}

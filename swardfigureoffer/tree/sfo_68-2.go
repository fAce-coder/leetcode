package tree

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 	递归

	/*
		p，q两个节点的最近公共祖先只有可能是以下几种情况：
			p，q分别在某个节点的左右子树中存在，则当前节点就是p，q的最近公共祖先
			p是p和q的公共祖先，q在p的左子树或者右子树中
			q是p和q的公共祖先，p在q的左子树或者右子树中
		因此，想要找到p和q的最近公共祖先，则只需要找到p或者q的位置即可
		1.终止条件：
			1.1 如果根节点是nil，则返回nil
			1.2 如果越过了叶子节点，还没有找到p或者q，则返回nil
			1.3 如果当前节点就是p或者q，则返回当前节点
		2.本次递归做的事：在当前节点为根节点的树中，找到p或者q，即判断当前节点是不是p或者q
		3.递归：递归地在当前节点的左右孩子为根节点的树中，找到p或者q，并判断p和q的位置
	*/

	// 1.终止条件
	// 1.1 如果根节点是nil，则返回nil
	// 1.2 如果找到为空的叶子节点，还没有找到p或者q，则返回nil
	if root == nil {
		return nil
	}
	// 1.3 如果当前节点就是p或者q，则返回当前节点
	if root == p || root == q {
		return root
	}

	// 2.递归：递归地在当前节点的左右孩子为根节点的树中，找到p或者q
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// 3.本次递归做的事：在当前节点为根节点的树中，找到p或者q,判断p和q的位置
	if left == nil && right == nil {
		// 3.1 如果左右子树都没有p和q，说明该树中根本没有p和q，则返回nil
		return nil
	} else if left != nil && right != nil {
		// 3.2 如果左右子树都找到了p和q，则说明p，q的位置分布在当前节点两侧，则当前节点就是最近公共祖先
		return root
	} else if left == nil && right != nil {
		// 3.3 如果左子树中没有p或者q，但是右子树中有，则说明此时right要么是最近公共祖先，要么是p或者q
		return right
	} else if left != nil && right == nil {
		// 3.4 如果右子树中没有p或者q，但是左子树中有，则说明此时left要么是最近公共祖先，要么是p或者q
		return left
	}

	return nil
}

package tree

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	// 递归

	/*
		判断B是否是A的子结构，则要尝试在A树中找到B树
		反过来说，如果B是A的子结构，则B的根节点一定是A中的某个节点
		因此，首先要在树A中找到B的根节点，然后判断在树A中以此节点为根的子树，是否跟树B相同，即：递归的判断其左右子树是否和B的左右子树相同
		本题分为两个步骤：
		1.第一步：找，先在A中找到跟B的根节点相同的节点；如果没找到，则在A递归的查找该节点的左右孩子，判断是否根B的根节点相同
		2.第二步：比，当在A中找到B的根节点后，就递归的比较A中那个节点的左右子树是否和B的根节点的左右子树相同；如果B都比较完了还相同，说明B是A的子树
	*/

	// 1.定义匿名函数，用来比较B是否是A的子树，即：比较根节点和后续节点是否相同
	var isSame func(*TreeNode, *TreeNode) bool
	isSame = func(nodeA *TreeNode, nodeB *TreeNode) bool {
		// 1.1 如果nodeB是空，说明此时B在这一支路已经匹配完，返回True
		if nodeB == nil {
			return true
		}
		// 1.2 如果nodeA是空，且nodeB不是空，说明此时nodeB还没有完全匹配，但nodeA已经没了，此时一定匹配不上，返回False
		if nodeA == nil {
			return false
		}
		// 1.3 如果nodeA和nodeB的值不相等，则匹配不上，返回false
		// 两种情况：第一次的根节点不相等；后续递归的左右孩子节点不相等
		if nodeA.Val != nodeB.Val {
			return false
		}
		// 1.4 递归的判断左右子树是否相等
		return isSame(nodeA.Left, nodeB.Left) && isSame(nodeA.Right, nodeB.Right)
	}

	// 2.如果A是空，或者B是空，则此时已经匹配不到了，返回False
	if A == nil || B == nil {
		return false
	}
	// 3.比，判断是否A的当前节点是B的根节点，且后续的左右子树都能包含B的左右子树
	if isSame(A, B) {
		return true
	}
	// 4.找，在A递归的查找该节点的左右孩子，判断是否以A的左右孩子为根节点与B的根节点做对比能确认B是A的子树
	return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

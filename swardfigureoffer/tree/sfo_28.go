package tree

func isSymmetric1(root *TreeNode) bool {
	// 递归

	/*
		新建递归函数实现比较两个节点的值是否相等
		终止条件：
			1.两个节点都是nil，说明都找到了最后的空叶子节点，且都是nil说明相等，返回True
			2.两个节点一个是nil，一个不是nil，说明其中一个找到最后的空叶子节点而另一个没有，则两个节点不相等，返回False
			3.两个节点都不是nil，且不相等，则返回False；如果相等，先不return，继续向下找，最终会找到终止条件1的状态才说明找完了，那时才返回True
		在本函数内要做的事：无
		递归：比较第一个节点的左孩子和第二个节点的右孩子，比较第一个节点的右孩子和第二个节点的左孩子
	*/

	// 1.特殊情况，当根节点是空，则直接返回true
	if root == nil {
		return true
	}

	// 2.新建一个匿名函数作为递归函数，实现比较两个节点的值是否相等
	var compareTwoNode func(*TreeNode, *TreeNode) bool
	compareTwoNode = func(left *TreeNode, right *TreeNode) bool {
		// 2.1 终止条件
		if left == nil && right == nil {
			// 2.1.1 两个节点都是nil，说明都找到了最后的空叶子节点，且都是nil说明相等，返回True
			return true
		} else if (left == nil && right != nil) || (left != nil && right == nil) {
			// 2.1.2 两个节点一个是nil，一个不是nil，说明其中一个找到最后的空叶子节点而另一个没有，则两个节点不相等，返回False
			return false
		} else if left.Val != right.Val {
			// 2.1.3 两个节点都不是nil，且不相等，则返回False；如果相等，先不return，继续向下找，最终会找到终止条件1的状态才说明找完了，那时才返回True
			return false
		}
		// 2.2 递归调用
		// 2.2.1 比较左节点的左孩子和右节点的右孩子
		res1 := compareTwoNode(left.Left, right.Right)
		// 2.2.2 比较左节点的右孩子和右节点的左孩子
		res2 := compareTwoNode(left.Right, right.Left)
		// 2.2.3 将递归比较的结果相与
		res := res1 && res2
		// 2.3 将结果返回
		return res
	}
	// 3.从根节点的左右孩子开始递归比较
	return compareTwoNode(root.Left, root.Right)
}

func isSymmetric2(root *TreeNode) bool {
	// 迭代

	/*
		使用队列来实现迭代
		将对称位置的两个节点作为一组数据压入队列中，出队后进行比较
		然后再分别将这两个节点的对称位置的左右孩子入队，以此类推，循环出队比较，直到队列为空，也就是所有节点都比较完毕
	*/

	// 1.特殊情况，当根节点是空，则直接返回true
	if root == nil {
		return true
	}
	// 2.初始化
	// 2.1 初始化一个用来存储一组节点数组的队列
	var treeNodeQueue [][2]*TreeNode = make([][2]*TreeNode, 0)
	// 2.2 初始化一个用来存储一组节点的数组，并实现复用
	// 注意：这里不能用切片，因为切片是引用类型，每次复用都会影响之前的数据
	var twoSymmetryNode [2]*TreeNode
	// 2.3 将根节点的左右孩子加入
	twoSymmetryNode[0] = root.Left
	twoSymmetryNode[1] = root.Right
	treeNodeQueue = append(treeNodeQueue, twoSymmetryNode)
	// 3.当队列不为空，开始迭代判断
	for len(treeNodeQueue) != 0 {
		// 3.1 出队
		left := treeNodeQueue[0][0]
		right := treeNodeQueue[0][1]
		treeNodeQueue = treeNodeQueue[1:]
		// 3.2 判断两个节点的状态
		if left == nil && right == nil {
			// 3.2.1 两个节点都是nil，说明都找到了最后的空叶子节点，且都是nil说明相等，跳过此次for循环
			continue
		} else if (left == nil && right != nil) || (left != nil && right == nil) {
			// 3.2.2 两个节点一个是nil，一个不是nil，说明其中一个找到最后的空叶子节点而另一个没有，则两个节点不相等，返回False
			return false
		} else if left.Val != right.Val {
			// 3.2.3 两个节点都不是nil，且不相等，则返回False
			return false
		}
		// 3.3 将这两个节点的对称位置的左右孩子入队
		twoSymmetryNode[0] = left.Left
		twoSymmetryNode[1] = right.Right
		treeNodeQueue = append(treeNodeQueue, twoSymmetryNode)

		twoSymmetryNode[0] = left.Right
		twoSymmetryNode[1] = right.Left
		treeNodeQueue = append(treeNodeQueue, twoSymmetryNode)

	}
	// 4.如果for循环结束还没有返回false，说明树对称，返回true
	return true

}

package tree

func kthLargest1(root *TreeNode, k int) int {
	// 递归

	/*
		二叉搜索树的中序遍历（左根右）是递增序列
		题目要求找出第k大的节点，即在中序遍历中倒数第k的节点
		所以要找到第k大的节点，就要按照中序遍历的反方向，即右根左遍历这棵树，得到的就是递减序列，即第k个节点就是第k大的节点
		终止条件：root = nil，说明根节点是空，或者找到了空的叶子节点，此时返回
		在本地递归中的操作：判断当前节点是不是第k大的节点，如果是，就返回这个结果
		递归：按照右根左的顺序递归查找这棵树
	*/

	// 1.定义要返回的结果
	var res int
	// 2.定义一个匿名函数作为递归函数
	var find func(*TreeNode)
	find = func(node *TreeNode) {
		// 2.1 终止条件：root = nil，说明根节点是空，或者找到了空的叶子节点，此时返回
		if node == nil {
			return
		}
		// 2.2 递归的查找右孩子
		find(node.Right)
		// 2.3 判断当前节点是不是第k大的节点，如果是，就返回这个结果
		k--
		if k == 0 {
			res = node.Val
			return
		}
		// 2.4 递归的查找左孩子
		find(node.Left)
	}
	// 3.从根节点开始查找
	find(root)
	// 4.返回结果
	return res
}

func kthLargest2(root *TreeNode, k int) int {
	// 颜色标记法

	/*
		使用栈来实现树的遍历
		二叉搜索树的中序遍历（左根右）是递增序列
		题目要求找出第k大的节点，即在中序遍历中倒数第k的节点
		所以要找到第k大的节点，就要按照中序遍历的反方向，即右根左遍历这棵树，得到的就是递减序列，即第k个节点就是第k大的节点
		要想得到右根左的遍历，就要按照左根右入栈，此时右孩子一直在栈顶
		对没有出过栈的标记为white，对出过栈的标记为gray
		因右孩子一直在栈顶，所以一直会查找右孩子（white），直到查找到最右（gray），此时回溯
	*/

	// 1.初始化
	// 1.1 建立一个形如[node,color]的数组
	type nodeAndColor [2]interface{}
	// 1.2 栈
	var stack []nodeAndColor
	// 1.3 返回结果
	var res int

	// 2.将根节点入栈
	stack = append(stack, nodeAndColor{root, "white"})

	// 3.当栈不为空，循环出栈
	for len(stack) != 0 {
		// 3.1 出栈
		node := stack[len(stack)-1][0].(*TreeNode)
		color := stack[len(stack)-1][1].(string)
		stack = stack[:len(stack)-1]
		// 3.2 如果节点是nil，则说明根节点是nil，或者找到了空的叶子节点，跳过
		if node == nil {
			continue
		}
		// 3.3 如果节点不是空
		if color == "white" {
			// 3.3.1 颜色是white，说明没出过栈
			// 左孩子入栈
			stack = append(stack, nodeAndColor{node.Left, "white"})
			// 当前节点入栈，但标记为灰色（已出过栈）
			stack = append(stack, nodeAndColor{node, "gray"})
			// 右孩子入栈
			stack = append(stack, nodeAndColor{node.Right, "white"})
		} else {
			// 3.3.2 颜色是gray，说明出过栈
			// 判断当前节点是否是第k大的节点，如果是，就生成结果并跳出循环
			k--
			if k == 0 {
				res = node.Val
				break
			}
		}
	}
	// 4.返回结果
	return res
}

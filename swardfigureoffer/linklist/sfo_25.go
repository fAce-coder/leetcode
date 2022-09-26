package linklist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	// 辅助链表

	/*
	   两个链表递增排序，合并后的链表仍然递增
	   比较两个链表中的节点，哪个小就把哪个接到新链表上
	*/
	// 0.特殊情况
	if l1 == nil && l2 == nil {
		return l1
	}
	// 1.新建一个链表头结点
	var resList *ListNode = &ListNode{}
	var res *ListNode = resList
	// 2.遍历两个链表
	// 2.1 当两个链表都还有,哪个小就把哪个接到结果集中
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			resList.Next = l1
			l1 = l1.Next
		} else {
			resList.Next = l2
			l2 = l2.Next
		}
		resList = resList.Next
	}
	// 2.2 当l1不为空，l2为空，则把l1接在后面
	if l1 != nil && l2 == nil {
		resList.Next = l1
	}
	// 2.3 当l2不为空，l1为空，则把l2接在后面
	if l2 != nil && l1 == nil {
		resList.Next = l2
	}
	// 将结果链表返回
	return res.Next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	// 递归

	// 终止条件
	// 如果l1为空，就返回l2
	if l1 == nil {
		return l2
	}
	// 如果l2为空，就返回l1
	if l2 == nil {
		return l1
	}

	// 递归
	var resList *ListNode
	if l1.Val <= l2.Val {
		resList = l1
		// resList = resList.Next
		resList.Next = mergeTwoLists2(l1.Next, l2)
	} else {
		resList = l2
		// resList = resList.Next
		resList.Next = mergeTwoLists2(l1, l2.Next)
	}

	return resList
}

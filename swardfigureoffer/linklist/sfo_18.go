package linklist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode1(head *ListNode, val int) *ListNode {
	// 头尾双指针

	// 特殊情况：如果删除的是第一个节点，直接返回链表的第二个节点
	if head.Val == val {
		return head.Next
	}
	// 初始化快慢指针
	var slow *ListNode = head
	var fast *ListNode = head.Next
	// 遍历链表，找到要删除的节点
	for fast != nil && fast.Val != val {
		slow = slow.Next
		fast = fast.Next
	}
	// 如果有要删除的节点，则删除节点
	if fast != nil {
		slow.Next = fast.Next
		fast.Next = nil
	}

	// 返回结果
	return head

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode2(head *ListNode, val int) *ListNode {
	// 递归

	// 终止条件
	if head.Val == val {
		return head.Next
	}

	// 递归
	head.Next = deleteNode2(head.Next, val)

	return head
}

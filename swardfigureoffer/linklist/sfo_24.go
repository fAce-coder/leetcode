package linklist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	// 迭代

	// 三指针
	var slow *ListNode
	var fast *ListNode = head
	// 逐渐迭代翻转
	for fast != nil {
		temp := fast.Next
		fast.Next = slow
		slow = fast
		fast = temp
	}
	// 此时slow慢指针的位置就是翻转后链表的头结点
	return slow
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList2(head *ListNode) *ListNode {
	// 递归

	// 终止条件
	if head == nil || head.Next == nil {
		return head
	}
	// 先递归到最后一个节点
	res := reverseList2(head.Next)
	// 从后向前处理当前节点
	head.Next.Next = head
	head.Next = nil

	return res
}

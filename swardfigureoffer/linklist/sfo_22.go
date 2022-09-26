package linklist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {
	// 快慢双指针

	/*
	   要找到链表倒数第k个节点，就让快指针先走k步
	   然后快慢指针一起走，直到快指针走到尾部，此时慢指针位置就是倒数第k个节点
	*/

	// 初始化快慢指针
	var slow *ListNode = head
	var fast *ListNode = head
	// 快指针先走k步
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	// 快慢指针同时移动，直到快指针走到尾部，慢指针处就是结果
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	// 返回慢指针
	return slow
}

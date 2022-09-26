package linklist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 双指针

	/*
	   A中不相遇的长度为a，B中不相遇的长度为b，相遇的长度为c
	   思路：A走到最后，从B开始走；B走到最后，从A开始走，相遇节点就是第一个公共节点
	   分以下两种情况：
	       当两个链表不一样长：a！=b
	           当两个链表不相交：AB都走了a + b + c的距离
	           当两个链表相交：A走了a + c + b，B走了b + c + a，相交在第一个公共节点
	       当两个链表一样长：a=b
	           当两个链表不相交：AB都走了a + c + b的距离
	           当两个链表相交：AB都走了a + c + b的距离，相交在第一个公共节点
	*/

	// 初始化双指针
	var nodeA *ListNode = headA
	var nodeB *ListNode = headB
	// 当nodeA与nodeB不相交，则遍历链表寻找
	for nodeA != nodeB {
		// A走到最后，从B开始走
		if nodeA != nil {
			nodeA = nodeA.Next
		} else {
			nodeA = headB
		}
		// B走到最后，从A开始走
		if nodeB != nil {
			nodeB = nodeB.Next
		} else {
			nodeB = headA
		}
	}
	// 返回相遇的点
	return nodeA
}

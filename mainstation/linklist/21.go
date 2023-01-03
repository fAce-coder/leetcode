package linklist

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 双指针

	/*
		因为两个链表都是升序的，若想把这两个链表合并成一个升序链表，只需分别在两个链表的首元结点放置1个指针
		新建一个头节点，作为返回结果的新链表的头节点
		比较两个指针所指向的节点值的大小，将较小的连接到新头节点后面，并将该指针向后移动一位；循环执行该步骤，直到两个指针都到达链表结尾
	*/

	// 1.特殊情况：
	// list1是空则直接返回list2
	if list1 == nil {
		return list2
	}
	// list2是空则直接返回list1
	if list2 == nil {
		return list1
	}

	// 2.新建一个头节点，作为返回结果的新链表res的头结点
	resHeader := &ListNode{}

	// 3.新建3个指针，分别指向list1、list2的首元结点和res的头结点
	p1 := list1
	p2 := list2
	pr := resHeader

	// 4.当p1和p2都不为空时，说明list1和list2都没比到结尾
	for p1 != nil && p2 != nil {
		// 此时进行比较，将较小节点的连接到res链表最后面，并将该指针向后移动一位以继续后续比较
		if p1.Val < p2.Val {
			pr.Next = p1
			p1 = p1.Next
		} else {
			pr.Next = p2
			p2 = p2.Next
		}
		// 将res链表指针向后移动一位，以继续连接节点
		pr = pr.Next
	}

	// 5.当p1比空了，说明list1比到了结尾，而list1和list2又是升序的，此时直接将list2的剩余部分接到res链表最后即可；p2比空后同理
	if p1 == nil {
		pr.Next = p2
	}
	if p2 == nil {
		pr.Next = p1
	}

	// 6.返回结果链表的首元结点（首元结点=头结点.Next）
	return resHeader.Next
}

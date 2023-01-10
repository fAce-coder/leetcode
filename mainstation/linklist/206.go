package linklist

func reverseList(head *ListNode) *ListNode {
	// 递归

	/*
		当前函数reverseList的作用是翻转以head作为首元节点的整个链表，并返回翻转后链表的头节点
		因此只需要直接使用此函数，获取以head.next作为首元节点的链表翻转后的头节点，相当于认定将head.next之后的节点都已经翻转好了
		最后再将head与head.next这两个节点之间翻转即可
	*/

	// 1.中止条件：
	// 		如果当前head是nil：链表长度为0
	// 		或者head.next是nil：链表长度为1，或无法继续递归获取head.next作为首元节点的链表翻转后的头节点
	if head == nil || head.Next == nil {
		return head
	}

	// 2.直接使用此函数，获取以head.next作为首元节点的链表翻转后的头节点
	// 相当于认定将head.next之后的节点都已经翻转好了，并返回了翻转之后链表的头节点
	newHead := reverseList(head.Next)

	// 3.最后再将head与head.next这两个节点之间翻转即可
	head.Next.Next = head
	head.Next = nil

	// 4.返回翻转后的新头节点
	return newHead
}

func reverseList1(head *ListNode) *ListNode {
	// 迭代：快慢指针

	/*
		1.起始时，将慢指针slow放在nil节点，即翻转链表后的尾节点的下一个节点；将快指针fast放在首元节点，即翻转后链表的尾节点
		2.获取临时指针temp来存储翻转前的fast.next节点
		3.执行翻转：将fast.next指向slow，将此处链表翻转
		4.slow和fast都向后移动一位，表示要翻转下一个位置，即：将slow更新为fast，将fast更新为temp
		5.直到fast走到了nil，即翻转前链表的尾节点的下一个节点，此时说明整个链表已经翻转完，此时slow处于翻转后链表的头结点，跳出循环
	*/

	// 1.特殊情况：链表中有0个或1个节点时，直接返回
	if head == nil || head.Next == nil {
		return head
	}

	// 2.初始化快慢指针
	var slow *ListNode // 慢指针放在nil节点，即翻转后链表的尾节点的下一个节点
	fast := head       // 将快指针fast放在首元节点，即翻转后链表的尾节点

	// 3.执行翻转链表
	for fast != nil {
		// 3.1 获取临时指针temp来存储翻转前的fast.next节点
		temp := fast.Next
		// 3.2 执行翻转：将fast.next指向slow，将此处链表翻转
		fast.Next = slow
		// 3.3 slow和fast都向后移动一位，表示要翻转下一个位置，即：将slow更新为fast，将fast更新为temp
		slow = fast
		fast = temp
	}

	// 4.直到fast走到了nil，即翻转前链表的尾节点的下一个节点，此时说明整个链表已经翻转完，此时slow处于翻转后链表的头结点，跳出循环，返回结果
	return slow
}

package linklist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint1(head *ListNode) []int {
	// 迭代法

	/*
	   定义三个指针，slow、fast、temp
	   逐渐向右移动，改变链表的指向
	*/

	var res []int = make([]int, 0)
	// 特殊情况，链表长度为0时
	if head == nil {
		return res
	}
	// 定义三指针
	var slow *ListNode
	var fast *ListNode = head
	var temp *ListNode
	// 遍历链表，翻转
	for fast != nil {
		temp = fast.Next
		fast.Next = slow
		slow = fast
		fast = temp
	}
	// 此时slow的位置就是翻转后新链表的起点
	for slow != nil {
		res = append(res, slow.Val)
		slow = slow.Next
	}
	// 返回结果
	return res
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint2(head *ListNode) []int {
	// 递归

	// 终止条件
	if head == nil {
		return []int{}
	}
	// 递归调用
	return append(reversePrint2(head.Next), head.Val)
}

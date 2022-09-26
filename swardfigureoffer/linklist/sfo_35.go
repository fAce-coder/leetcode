package linklist

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
func copyRandomList(head *Node) *Node {
	// 哈希表

	/*
	   建立一个哈希表，键值都是*Node类型
	       键：原节点
	       值：新建一个值与原节点相同的节点，并与键中的节点一一对应
	   遍历哈希表，hashMap[key].Next/Random = hashMap[key.Next/Random]
	   最后返回键为头结点对应的值，即为新复制的链表
	*/

	// 1.新建一个哈希表
	var hashMap map[*Node]*Node = make(map[*Node]*Node)
	// 2.遍历原链表，将原链表的节点作为键，新复制的节点作为值，添加到hashMap中
	var temp *Node = head
	for temp != nil {
		hashMap[temp] = &Node{
			Val: temp.Val,
		}
		temp = temp.Next
	}
	// 3.遍历哈希表，按照原链表，建立新链表的next和random关系
	for key, _ := range hashMap {
		// 3.1 构建Next
		if key.Next != nil {
			hashMap[key].Next = hashMap[key.Next]
		}
		// 3.2 构建Random
		hashMap[key].Random = hashMap[key.Random]
	}
	// 4.返回键为头节点所对应的值，即为新复制链表的头结点
	return hashMap[head]
}

package array

import (
	"container/list"
)

// CQueue 双栈实现队列
// 队列的数据结构
type CQueue struct {
	// 定义2个栈
	StackA []int
	StackB []int
}

// CQueueConstructor 构造函数
func CQueueConstructor() CQueue {
	// 返回一个队列对象
	return CQueue{
		StackA: make([]int, 0),
		StackB: make([]int, 0),
	}
}

// AppendTail 入队
func (q *CQueue) AppendTail(value int) {
	// 只往A中入栈
	q.StackA = append(q.StackA, value)
}

// DeleteHead 出队
func (q *CQueue) DeleteHead() int {
	// 1. A,B都空
	if len(q.StackA) == 0 && len(q.StackB) == 0 {
		return -1
		// 2. A不空，B空
	} else if len(q.StackA) != 0 && len(q.StackB) == 0 {
		// A循环出栈，并全部倒进B中
		for len(q.StackA) != 0 {
			topOfStackA := q.StackA[len(q.StackA)-1] // 获取A栈顶元素
			q.StackA = q.StackA[:len(q.StackA)-1]    // A出栈
			q.StackB = append(q.StackB, topOfStackA) // B入栈
		}
		// 从B中出栈
		topOfStackB := q.StackB[len(q.StackB)-1] // 获取B栈顶元素
		q.StackB = q.StackB[:len(q.StackB)-1]    // B出栈
		return topOfStackB                       // 将B栈顶元素作为结果返回
		// 3. A空B不空，A不空B不空 --> B不空
	} else if len(q.StackB) != 0 {
		// 从B中出栈
		topOfStackB := q.StackB[len(q.StackB)-1] // 获取B栈顶元素
		q.StackB = q.StackB[:len(q.StackB)-1]    // B出栈
		return topOfStackB                       // 将B栈顶元素作为结果返回
	}
	return -1
}

// CQueueStandard 标准库版
type CQueueStandard struct {
	list1, list2 *list.List // list包的List函数，是一个结构体
}

func CQueueStandardConstructor() CQueueStandard {
	return CQueueStandard{
		list1: list.New(), // New函数，返回*List对象，即创建链表
		list2: list.New(),
	}
}

func (qs *CQueueStandard) AppendTailStandard(value int) {
	qs.list1.PushBack(value) // 直接插入list1尾部即可
}

func (qs *CQueueStandard) DeleteHeadStandard() int {
	if qs.list2.Len() == 0 { // 如果list2为空就把list1的元素全部倒置放入list2中并在list1中删除
		for qs.list1.Len() != 0 {
			qs.list2.PushBack(qs.list1.Remove(qs.list1.Back()))
		}
	} // 这里看似上下两个if刚好条件是相反的，但是一定不能写为if-else，因为我们要的是顺序执行而不是二选一

	if qs.list2.Len() != 0 { // 不为空就返回list2的Back并删除
		e := qs.list2.Back()
		qs.list2.Remove(e)
		return e.Value.(int) // 注意V大写，并使用类型断言转为int型
	}
	return -1 // 没有则返回默认值-1

}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

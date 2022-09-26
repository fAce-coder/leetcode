package array

type MinStack struct {
	MinStack []int // 最小栈，栈顶维持着普通栈中的最小元素
	Stack    []int // 普通栈，正常入栈出栈
}

// MinStackConstructor /** initialize your data structure here. */
// 构造函数
func MinStackConstructor() MinStack {
	return MinStack{
		MinStack: make([]int, 0),
		Stack:    make([]int, 0),
	}
}

// Push 入栈
func (m *MinStack) Push(x int) {
	// 普通栈正常入栈
	m.Stack = append(m.Stack, x)
	// 最小栈入栈
	// 只有最小栈为空，或者要入栈的元素比栈顶元素小或等，才入栈
	if len(m.MinStack) == 0 || x <= m.MinStack[len(m.MinStack)-1] {
		m.MinStack = append(m.MinStack, x)
	}
}

// Pop 出栈
func (m *MinStack) Pop() {
	// 普通栈正常出栈
	popOfStack := m.Stack[len(m.Stack)-1] // 获取普通栈要出栈的元素
	m.Stack = m.Stack[:len(m.Stack)-1]    // 普通栈出栈
	// 只有当最小栈的栈顶元素=普通栈的出栈元素，此时最小栈才出栈
	if m.MinStack[len(m.MinStack)-1] == popOfStack {
		m.MinStack = m.MinStack[:len(m.MinStack)-1]
	}
}

// Top 获取栈顶元素
func (m *MinStack) Top() int {
	// 返回普通栈的栈顶元素
	return m.Stack[len(m.Stack)-1]
}

// Min 获取最小元素
func (m *MinStack) Min() int {
	// 返回最小栈的栈顶元素
	return m.MinStack[len(m.MinStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */

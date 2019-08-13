package stack

//定义栈结构
type Stack struct {
	node int     //当前的节点
	data [10]int //实际数据
}

//往栈中增加数据
func (stack *Stack) Push(number int) {
	if stack.node < len(stack.data) {
		stack.data[stack.node] = number
		stack.node++
	}

}

//从栈中取出数据
func (stack *Stack) Pop() int {
	if stack.node > -1 {
		stack.node--

	}
	return stack.data[stack.node]
}

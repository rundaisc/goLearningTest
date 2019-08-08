package main

import (
	"fmt"
	"strconv"
)

func main() {

}

//SortNumber 编写函数，返回其（两个）参数正确的（自然）数字顺序：
func SortNumber(a, b int) (int, int) {
	if a > b {
		return a, b
	}
	return b, a

}

// FoundError 下面的程序有什么错误？
func FoundError() {

	for i := 0; i < 10; i++ {
		fmt.Printf("%v\n", i)
	}
	// fmt.Printf("%v\n", i) //作用域错误。这里会提示I未定义
}

//定义栈结构
type stack struct {
	node int     //当前的节点
	data [10]int //实际数据
}

//往栈中增加数据
func (stack *stack) Push(number int) {
	if stack.node < len(stack.data) {
		stack.data[stack.node] = number
		stack.node++
	}

}

//从栈中取出数据
func (stack *stack) Pop() int {
	if stack.node > -1 {
		stack.node--

	}
	return stack.data[stack.node]
}

// 更进一步。编写一个 String 方法将栈转化为字符串形式的表达。可以这样的
//方式打印整个栈： fmt.Printf("My stack %v\n", stack)
//栈可以被输出成这样的形式： [0:m] [1:l] [2:k]
func (stack *stack) ToString() string {
	var s string
	for i := 0; i < stack.node; i++ {
		k := strconv.Itoa(i)
		v := strconv.Itoa(stack.data[i])
		s += "[" + k + ":" + v + "] "

	}

	return s
}

//Tooargs 编写函数接受整数类型变参，并且每行打印一个数字。
func Tooargs(args ...int) {
	for _, v := range args {
		fmt.Println(v)
	}

}

//Gblq 斐波那契数列以：1,1,2,3,5,8,13,... 开始
func Gblq(number int) []int {
	s := make([]int, number)
	for i := 1; i <= number; i++ {
		if i < 3 {
			s[i-1] = 1
		} else {
			s[i-1] = s[i-2] + s[i-3]

		}

	}
	return s

}

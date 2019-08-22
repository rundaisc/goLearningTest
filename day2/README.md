#### day2
- 编写函数，返回其（两个）参数正确的（自然）数字顺序：
- 创建一个固定大小保存整数的栈。它无须超出限制的增长。定义 push 函数——
  将数据放入栈，和 pop 函数——从栈中取得内容。栈应当是后进先出（LIFO）
  的。
  - 更进一步。编写一个 String 方法将栈转化为字符串形式的表达。可以这样的
    方式打印整个栈： fmt.Printf("My stack %v\n", stack)
    栈可以被输出成这样的形式： [0:m] [1:l] [2:k]
- 编写函数接受整数类型变参，并且每行打印一个数字。
-  斐波那契数列以：1,1,2,3,5,8,13,... 开始。或者用数学形式表达：x 1 = 1;x 2 =
  1;x n = x n−1 + x n−2 ∀n > 2。
  编写一个接受 int 值的函数，并给出这个值得到的斐波那契数列。
 - (1) map 函数 map() 函数是一个接受一个函数和一个列表作为参数的函数。函
   数应用于列表中的每个元素，而一个新的包含有计算结果的列表被返回。因此：
   map(f(),(a 1 ,a 2 ,...,a n−1 ,a n )) = (f(a 1 ),f(a 2 ),...,f(a n−1 ),f(a n ))
   1. 编写 Go 中的简单的 map() 函数。它能工作于操作整数的函数就可以了。
   2. 扩展代码使其工作于字符串列表。
 - )最小值和最大值
   1. 编写一个函数，找到 int slice ( []int ) 中的最大值。
   2. 编写一个函数，找到 int slice ( []int ) 中的最小值。
-   编写一个针对 int 类型的 slice 冒泡排序的函数
- 函数返回一个函数
    1. 编写一个函数返回另一个函数，返回的函数的作用是对一个整数 +2。函数的名
  称叫做 plusTwo 。然后可以像下面这样使用：
  p := plusTwo()
  fmt.Printf("%v\n", p(2))
  应该打印 4。参阅第 31 页的 “回调” 小节了解更多相关信息。
  2. 使 1 中的函数更加通用化，创建一个 plusX(x) 函数，返回一个函数用于对整
  数加上 x 。
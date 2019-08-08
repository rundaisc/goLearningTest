package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	First()
	Second()
	Third()
}

//First  创建一个基于 for 的简单的循环。使其循环 10 次，并且使用 fmt 包打印出计数器的值。
func First() {
	fmt.Println("first func start")
	for i := 1; i < 11; i++ {
		fmt.Println(i)
	}
	fmt.Println("first func end")
}

//Second  用 goto 改写 1 的循环。关键字 for 不可使用。
func Second() {
	fmt.Println("second func start")
	i := 1
HERE:
	fmt.Println(i)
	if i < 10 {
		i++
		goto HERE
	}
	fmt.Println("second func end")
}

//Third  再次改写这个循环，使其遍历一个 array，并将这个 array 打印到屏幕上。
func Third() {
	fmt.Println("third func start")
	arr := [5]int{1, 2, 3, 4, 5}
	for i, v := range arr {
		fmt.Printf("第%d个元素为%d\n", i+1, v)
	}
	fmt.Println("third func end")
}

//FizzBuzz  编写一个程序，打印从 1 到 100 的数字。当是三的倍数就打印 “Fizz”代替数字，
//当是5的倍数就打印 “Buzz” 。当数字同时是三和5的倍数时，打印 “FizzBuzz” 。
func FizzBuzz(number int) {
	fmt.Println("FizzBuzz func start")
	for i := 1; i <= number; i++ {
		switch true {

		case i%3 == 0 && i%5 == 0:
			fmt.Printf("%s ", "FizzBuzz")
		case i%3 == 0:
			fmt.Printf("%s ", "Fizz")
		case i%5 == 0:
			fmt.Printf("%s ", "Buzz")
		default:
			fmt.Printf("%d ", i)
		}

	}
	fmt.Println()
	fmt.Println("FizzBuzz func end")

}

//PrintA 建立一个 Go 程序打印下面的内容（到 100 个字符）：  A
// AA
// AAA
// AAAA
// AAAAA
// AAAAAA
// AAAAAAA
func PrintA(number int) {
	a := "A"
	for i := 0; i < number; i++ {
		fmt.Println(a)
		a = a + "A"
	}
}

//CountStr 建立一个程序统计字符串里的字符数量：
// asSASA ddd dsjkdsjs dk你好
// 同时输出这个字符串的字节数。
func CountStr() {
	str := "asSASA ddd dsjkdsjs dk你好"
	strRune := []rune(str)
	fmt.Printf("len长度为:%d\n", len(str))
	fmt.Printf("rune长度为:%d\n", len(strRune))
	fmt.Printf("RuneCountInString长度为:%d\n", utf8.RuneCountInString(str))
}

//ChangeStr 扩展/修改上一个问题的程序，替换位置 4 开始的三个字符为 “abc”。
func ChangeStr() {
	str := "asSASA ddd dsjkdsjs dk你好"
	strRune := []rune(str)
	//第一种
	// strRune[4] = 'a'
	// strRune[5] = 'b'
	// strRune[6] = 'c'

	//第二种
	copy(strRune[4:4+3], []rune("abc"))
	cStr := string(strRune)

	fmt.Println(cStr)
}

//ReverseStr  编写一个 Go 程序可以逆转字符串，例如 “foobar” 被打印成 “raboof”
func ReverseStr(str string) {
	rs := []rune(str)
	// 第一种
	var newStr string
	for i := len(rs) - 1; i >= 0; i-- {

		newStr = newStr + string(rs[i])
	}
	fmt.Println(newStr)
	// 第二种
	var newStr2 []rune

	for i := len(rs) - 1; i >= 0; i-- {
		newStr2 = append(newStr2, rs[i])
	}
	fmt.Println(string(newStr2))

}

// AverageSlice 编写计算一个类型是 float64 的 slice 的平均值的代码。
func AverageSlice() {
	s := []float64{103.44, 20.07, 14.11, 1.01}
	var total float64
	for _, v := range s {
		total += v
	}
	fmt.Println(total / float64(len(s)))
}

package main

import "fmt"

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

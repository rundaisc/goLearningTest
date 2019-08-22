package day4

import (
	"fmt"
)

//First  创建一个基于 for 的简单的循环。使其循环 10 次，并且使用 fmt 包打印出计数器的值。channel
func First(){
	/*
	利用 sync 包

	var wg sync.WaitGroup
	wg.Add(10)
	for i:=1 ;i<11;i++{
		go func(number int) {
			fmt.Println(number)
		}(i)
	wg.Done()
	}
	wg.Wait()
	*/

	//利用channel
	c := make(chan int)

	flag := make ( chan bool ) //用户判断是否执行完成,如果没有这个判断，那么输出结果 只会到9
	go func(ch chan int,f chan bool) {
		for  {
			select {
			case <-flag :
				break
			case number := <-ch:
				fmt.Println(number)
			}
		}
	}(c,flag)

	for i:=1;i<11;i++{
		c<-i
	}

	flag<-false

	//close(c)

}

func  dup3(in <-chan  int) (<-chan int,<-chan int,<-chan int)  {
	a,b,c := make(chan int ,2) ,make(chan int ,2),make(chan int ,2)
	go func() {
		for  {
			x := <-in
			a<-x
			b<-x
			c<-x
		}
	}()
	return  a,b,c
}

func fib() <-chan int{
	x := make(chan int,2)
	a ,b ,c := dup3(x)
	go func() {
		x<-0
		x<-1
		<-a
		for  {
			x<- <-a+<-b
		}
	}()
	return  c
}
//实现斐波拉契 使用channel
func Runfib()  {
	ch := fib()
	for i:=0;i<20 ; i++ {
		fmt.Println(<-ch)
	}
}



package day5

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
	for i:=1;i<11;i++{
		go func(number int) {
			fmt.Printf("%d from goroutine\r\n",number)
			c<-i
		}(i)
		i,_:= <-c
		fmt.Printf("%d from channel\r\n",i)
	}
	close(c)

}

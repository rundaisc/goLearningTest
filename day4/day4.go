package day4

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type base interface {}

//通过typeSwitch 判断类型
func typeSwitch(f base) base {
	switch f.(type) {
	case int:
		return  f.(int)*2;
	case string:
		return f.(string)+f.(string)+f.(string)
	}
	return  f
}
//对数组进行迭代操作
func Map(arr []base,f func( base) base)[]base{
	for k,v :=range arr{
		arr[k] = f(v)
	}

	return  arr
}
//模拟 cat 读取文件
func Cat(fileName string,showLine bool){
	file,err := os.OpenFile(fileName,os.O_RDWR,0666)
	if err!=nil{
		fmt.Println("open file err",err)
		return
	}
	defer file.Close()
	buf := bufio.NewReader(file)
	lineNumber := 1
	for {
		line,err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if showLine{
			fmt.Println(strconv.Itoa(lineNumber)+" "+line)
		}else {
			fmt.Println(line)
		}
		lineNumber++
		if err != nil{
			if err == io.EOF {
				break
			}else {
				fmt.Println("open file err",err)
				return
			}

		}
	}





}





package day3

import (
	"testing"
)

func TestMap(t *testing.T) {
	numbers := []base{1,2,3,4}
	strings := []base{"2","b","c"}
	newNumbers := Map(numbers,typeSwitch)
	newStrings := Map(strings,typeSwitch)
	t.Log(newNumbers)
	t.Log(numbers) //说明是引用类型
	t.Log(newStrings)
	t.Log(strings) //说明是引用类型
}

func  TestCat(t *testing.T) {
	Cat("log.txt",true)
}

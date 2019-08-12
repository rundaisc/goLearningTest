package main

import (
	"testing"
)

func TestSortNumber(t *testing.T) {
	a, b := SortNumber(7, 2)
	t.Log(a, b)

}

func TestPush(t *testing.T) {
	s := stack{}
	for index := 1; index < 11; index++ {
		s.Push(index)
	}
	t.Log(s)
	for index := 1; index < 11; index++ {
		t.Log(s.Pop())

	}

}

func TestToString(t *testing.T) {
	s := stack{}
	for index := 1; index < 11; index++ {
		s.Push(index)
	}
	t.Log(s.ToString())

}

func TestTooargs(t *testing.T) {
	Tooargs(1, 2, 3, 4)
}

func TestGblq(t *testing.T) {
	s := Gblq(20)
	t.Log(s)
}

func TestArrayMap(t *testing.T) {
	arr := []int{1, 3, 4}
	mapArr := ArrayMap(arr, func(arg1 int) int {
		return arg1 * arg1
	})
	t.Log(mapArr)

}

func TestMaxOf(t *testing.T) {
	arr := []int{9, 3, 4, 7, 11, 2}
	max := MaxOf(arr)
	t.Log(max)
}

func TestMinOf(t *testing.T) {
	arr := []int{9, 3, 4, 7, 11, 2}
	min := MinOf(arr)
	t.Log(min)
}

func TestArrSort(t *testing.T) {
	arr := []int{9, 3, 4, 7, 11, 2}
	ArrSort(arr)
	t.Log(arr)
}

func TestPlusTwo(t *testing.T) {
	t.Log(PlusTwo()(3))
	//或者
	f := PlusTwo()
	t.Log(f(3))
}

func TestPlusX(t *testing.T) {

	f := PlusX(7)
	t.Log(f(3))
}

package main

import (
	"fmt"
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
	fmt.Println(s)
	for index := 1; index < 11; index++ {
		fmt.Println(s.Pop())
	}

}

func TestToString(t *testing.T) {
	s := stack{}
	for index := 1; index < 11; index++ {
		s.Push(index)
	}
	fmt.Println(s.ToString())

}

func TestTooargs(t *testing.T) {
	Tooargs(1, 2, 3, 4)
}

func TestGblq(t *testing.T) {
	s := Gblq(20)
	t.Log(s)
}

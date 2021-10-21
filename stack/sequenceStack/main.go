package main

import (
	"fmt"
)

// 使用数组模拟栈
type Stack struct {
	// 栈的容量
	MaxTop int
	// 栈堆
	Top int
	// 数组模拟栈
	Array [5]int
}

func (s *Stack) Push(val int) {
	// fmt.Println(len(s.Array))
	// 判断栈是否满了
	if s.Top >= s.MaxTop-1 {
		fmt.Println("the stack is full")
		return
	}

	s.Top++
	// 入栈
	s.Array[s.Top] = val
	return
}

// 出栈
func (s *Stack) Pop() (data int) {
	// 判断栈是否为空
	if s.Top == -1 {
		fmt.Println("the stack is empty")
		return
	}

	data = s.Array[s.Top]
	s.Top--
	return
}

// 遍历栈
func (s *Stack) List() {
	// 判断栈是否满了
	if s.Top == -1 {
		fmt.Println("the stack is empty")
		return
	}
	for i := s.Top; i >= 0; i-- {
		fmt.Println(s.Array[i])
	}
}

// 初始化一个栈
func InitStack(maxTop int) *Stack {
	return &Stack{
		MaxTop: maxTop,
		Top:    -1,
	}
}
func main() {
	stack := InitStack(5)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.List()
}

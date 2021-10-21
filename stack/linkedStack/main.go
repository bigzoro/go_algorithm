package main

import (
	"errors"
	"fmt"
)

// 栈节点
type Node struct {
	data int
	next *Node
}

// 栈链表
type Stack struct {
	maxLength int
	length    int
	head      *Node
}

// 初始化一个栈
func InitStack(maxLength int) *Stack {
	return &Stack{
		maxLength: maxLength,
		length:    0,
		head:      nil,
	}
}

// 入栈
func (s *Stack) Push(val int) {

	// 判断栈是否满了
	if s.length >= s.maxLength {
		fmt.Println("the stack is full")
		return
	}

	// 生成要入栈的节点
	node := &Node{
		data: val,
		next: s.head,
	}
	// 入栈
	s.head = node
	s.length++
}

// 取出一个元素
func (s *Stack) Pop() (data int, err error) {
	if s.length <= 0 {
		err = errors.New("the stack is empty")
		return
	}
	data = s.head.data
	// 指向后面的一个节点
	s.head = s.head.next
	s.length--
	return
}

// 查看所有元素
func (s *Stack) List() {
	// 判断是否为空
	if s.length == 0 {
		fmt.Println("the stack is empty")
		return
	}

	temp := s.head
	for i := 0; i < s.length; i++ {
		fmt.Println(temp.data)
		temp = temp.next
	}
}

func main() {
	stack := InitStack(5)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.List()
	data, err := stack.Pop()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)

}

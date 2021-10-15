package main

import (
	"errors"
	"fmt"
	"os"
)

// 队列节点
type Node struct {
	element int
	next    *Node
}

// 队列链表
type LinkedQueue struct {
	Front  *Node
	Rear   *Node
	length int
}

// 初始化队列
func (lq *LinkedQueue) initLinkedQueue() {
	lq.Front = nil
	lq.Rear = nil
	lq.length = 0
	fmt.Printf("初始化成功！\n")
}

// 判断队列是否为空
func (lq *LinkedQueue) isEmpty() bool {
	flag := false
	if lq.length == 0 {
		flag = true
	}
	return flag
}

// 入队
func (lq *LinkedQueue) add(data int) {

	// 构造新的节点
	node := Node{element: data}

	// 判断是否为第一次插入
	if lq.isEmpty() {
		lq.Front = &node
		lq.Rear = &node
	} else {
		lq.Rear.next = &node
		lq.Rear = &node
	}
	lq.length++
}

// 出队
func (lq *LinkedQueue) poll() (data int, err error) {
	if lq.isEmpty() {
		err = errors.New("the quenen is empty")
		return
	}

	// 取出节点的数据
	data = lq.Front.element

	// 处理剩下最后一个节点的情况
	if lq.length == 1 {
		lq.Front = nil
		lq.Rear = nil
		lq.length--
		return
	}

	lq.Front = lq.Front.next
	lq.length--

	return data, err
}

// 查看队列的所有元素
func (lq *LinkedQueue) list() {
	if lq.isEmpty() {
		fmt.Println("the queue is empty")
		return
	}
	// 因为队列头不能变，所有用临时变量代替一下
	tempFront := lq.Front
	for i := 0; i < lq.length; i++ {
		fmt.Println(tempFront.element)
		tempFront = tempFront.next
	}
}
func main() {
	linkedQueue := LinkedQueue{}
	var key int
	var value int
	for {

		fmt.Println("0. 初始化链式队列")
		fmt.Println("1. 添加数据到链式队列")
		fmt.Println("2. 从链式队列中获取数据")
		fmt.Println("3. 显示链式队列数据")
		fmt.Println("4. 退出程序")
		fmt.Println("请选择（0-4）：")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 0:
			linkedQueue.initLinkedQueue()
		case 1:
			fmt.Println("请输入要添加的数据：")
			fmt.Scanf("%d\n", &value)
			linkedQueue.add(value)
		case 2:
			data, err := linkedQueue.poll()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("从队列中取出的数据为： %d\n", data)
		case 3:
			linkedQueue.list()
		case 4:
			os.Exit(0)
		}
	}
}

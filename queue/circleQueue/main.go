package main

import (
	"errors"
	"fmt"
	"os"
)

type CircleQueue struct {
	MaxSize int
	Front   int
	Rear    int
	Element []int
}

// 初始化队列
func (cq *CircleQueue) initCircleQueue(maxSize int) {
	cq.MaxSize = maxSize
	cq.Element = make([]int, maxSize)
	cq.Front = 0
	cq.Rear = 0
}

// 判断队列是否为空
func (cq *CircleQueue) isEmpty() bool {
	flag := false
	if cq.Front == cq.Rear {
		flag = true
	}
	return flag
}

// 判断队列是否已满
func (cq *CircleQueue) isFull() bool {
	flag := false
	if (cq.Rear+1)%cq.MaxSize == cq.Front {
		flag = true
	}
	return flag
}

// 插入一个数据
func (cq *CircleQueue) add(data int) (err error) {
	if cq.isFull() {
		err = errors.New("the queue is full")
		return
	}
	cq.Element[cq.Rear] = data
	cq.Rear = (cq.Rear + 1) % cq.MaxSize
	return
}

// 取出一个数据
func (cq *CircleQueue) poll() (data int, err error) {
	if cq.isEmpty() {
		err = errors.New("the queue is empty")
		return
	}
	data = cq.Element[cq.Front]
	cq.Front = (cq.Front + 1) % cq.MaxSize
	return data, err
}

// 显示队列中所有的值
func (cq *CircleQueue) list() {
	if cq.isEmpty() {
		fmt.Println("the queue is empty")
		return
	}
	// 因为队列的头不能改变，所以设置一个辅助遍历
	tempFront := cq.Front
	for i := 0; i < (cq.Rear+cq.MaxSize-cq.Front)%cq.MaxSize; i++ {
		fmt.Printf("arr[%d]=%d\n", tempFront, cq.Element[tempFront])
		tempFront = (tempFront + 1) % cq.MaxSize
	}
}
func main() {

	circleQueue := CircleQueue{}
	var key int
	var value int
	var size int
	for {

		fmt.Println("0. 初始化环形队列")
		fmt.Println("1. 添加数据到环形队列")
		fmt.Println("2. 从环形队列中获取数据")
		fmt.Println("3. 显示环形队列数据")
		fmt.Println("4. 退出程序")
		fmt.Println("请选择（0-4）：")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 0:
			fmt.Println("请输入初始化队列的容量")
			fmt.Scanf("%d\n", &size)
			circleQueue.initCircleQueue(size)
		case 1:
			fmt.Println("请输入要添加的数据：")
			fmt.Scanf("%d\n", &value)
			err := circleQueue.add(value)
			if err != nil {
				fmt.Println(err)
				return
			}
		case 2:
			data, err := circleQueue.poll()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("从队列中取出的数据为： %d\n", data)
		case 3:
			circleQueue.list()
		case 4:
			os.Exit(0)
		}
	}
}

package main

import (
	"errors"
	"fmt"
	"os"
)

type Queue struct {
	MaxSize int
	Front   int
	Rear    int
	Element []int
}

// 初始化队列
func (q *Queue) initQueue(maxSize int) {
	q.MaxSize = maxSize
	q.Element = make([]int, maxSize)
	q.Front = -1
	q.Rear = -1
}

// 判断队列是否为空
func (q *Queue) isEmpty() bool {
	flag := false
	if q.Front == q.Rear {
		flag = true
	}
	return flag
}

// 向队列中添加一个元素
func (q *Queue) add(data int) (err error) {
	// 判断满了
	if q.Rear == q.MaxSize-1 {
		err = errors.New("队列已满，无法添加！")
		return err
	}
	q.Rear++
	q.Element[q.Rear] = data
	return
}

// 获取队列的一个元素
func (q *Queue) poll() (data int, err error) {
	if q.isEmpty() {
		err = errors.New("the queue is empty")
		return
	}
	q.Front++
	data = q.Element[q.Front]
	return data, err
}

// 查看队列的所有元素
func (q *Queue) list() {
	for i := q.Front + 1; i < len(q.Element)-1; i++ {
		fmt.Printf("element[%d]：%d\n", i, q.Element[i])
	}
}
func main() {
	queue := Queue{}
	var key int
	var value int
	var size int
	for {

		fmt.Println("0. 初始化队列")
		fmt.Println("1. 添加数据到队列")
		fmt.Println("2. 从队列中获取数据")
		fmt.Println("3. 显示队列数据")
		fmt.Println("4. 退出程序")
		fmt.Println("请选择（0-4）：")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 0:
			fmt.Println("请输入初始化队列的容量")
			fmt.Scanf("%d\n", &size)
			queue.initQueue(size)
		case 1:
			fmt.Println("请输入要添加的数据：")
			fmt.Scanf("%d\n", &value)
			err := queue.add(value)
			if err != nil {
				fmt.Println(err)
				return
			}
		case 2:
			data, err := queue.poll()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("从队列中取出的数据为： %d\n", data)
		case 3:
			queue.list()
		case 4:
			os.Exit(0)
		}
	}
}

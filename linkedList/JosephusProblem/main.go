package main

import "fmt"

/*
	问题描述：设标号为1，2..., n的n个人围坐在一圈，约定编号为k(1<=k<=n)的人从1开始报数，数到m的那个人出列，
		它的下一位又开始从1开始报数，数到m的那个人又出列，以此类推，直到所有人出列为止，由此产生一个出队编号的序列
*/

// 小孩的结构体
type Boy struct {
	No   int
	next *Boy
}

// num：表示小孩的个数
// *Boy：返回该环形链表的第一个小孩的指针
func AddBoy(num int) *Boy {
	first := &Boy{}
	currentBoy := &Boy{} // 辅助指针
	if num < 1 {
		fmt.Println("number is uncorrect")
		return first
	}

	// 循环构建链表
	for i := 1; i <= num; i++ {
		boy := &Boy{
			No: i,
		}
		// 第一个小孩
		if i == 1 {
			// 第一个小孩的头指针不能动，否则后面的节点就无法找到第一个小孩的位置
			first = boy
			currentBoy = boy
			currentBoy.next = first // 形成循环
		} else {
			currentBoy.next = boy
			currentBoy = boy
			currentBoy.next = first
		}
	}
	return first
}

// 显示单向的环形链表
func List(first *Boy) {
	// 处理空的环形链表
	if first.next == nil {
		fmt.Println("链表为空")
		return
	}

	currentBoy := first
	for {
		fmt.Printf("小孩编号：%d\n", currentBoy.No)
		// 退出条件
		if currentBoy.next == first {
			break
		}
		currentBoy = currentBoy.next
	}
}

/*
	编写一个函数，palyGame(first *Boy, startNo int, countNum int)
	使用算法，按照要求，在环形链表中留下最后一个人
*/
func palyGame(first *Boy, startNo int, countNum int) {

	// 空链表处理
	if first.next == nil {
		fmt.Println("null")
		return
	}
	// 辅助指针，帮助删除小孩
	tail := first
	// 让tail指向环形链表的最后一个节点`
	for {
		// 到了最后一个小孩了
		if tail.next == first {
			break
		}
		tail = tail.next
	}
	// 让first移动到startNo
	for i := 0; i < startNo-1; i++ {
		first = first.next
		tail = tail.next
	}

	// 开始数countNum下，然后删除first指向的小孩
	for {
		// 开始数countNUm-1下，自己也要数
		for i := 0; i < countNum-1; i++ {
			first = first.next
			tail = tail.next
		}
		fmt.Println("出圈小孩为：", first.No)
		// 删除first指向的节点
		first = first.next
		tail.next = first

		// 剩下最后一个小孩不删除
		if tail == first {
			break
		}
	}
	fmt.Printf("%d小孩编号\n", first.No)
}
func main() {
	first := AddBoy(5)
	palyGame(first, 2, 3)
}

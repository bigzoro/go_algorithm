package main

import (
	"fmt"
	"os"
)

// 员工节点
type Emp struct {
	Id   int
	Name string
	Next *Emp
}

// 展示员工信息的方法
func (e *Emp) ShowMe() {
	fmt.Printf("链表%d 找到了雇员%d\n", e.Id%7, e.Id)
}

// 定义员工链表，这里的员工链表不带表头，即第一个节点就存放雇员
type EmpLink struct {
	Head *Emp
}

// 添加员工，添加时从小到大
func (el *EmpLink) Insert(emp *Emp) {
	// 循环找到合适的位置，进行添加
	temp := el.Head
	if temp == nil {
		el.Head = emp
		return
	}
	for {
		if temp.Next == nil {
			break
		}
		if temp.Next.Id > emp.Id {
			break
		}
		temp = temp.Next
	}
	emp.Next = temp.Next
	temp.Next = emp

}

// 通过员工ID查找员工
func (el *EmpLink) FindById(id int) (emp *Emp) {
	temp := el.Head
	flag := false
	if temp == nil {
		fmt.Println("the empLink is empty")
		return
	}
	for {
		// 找到了员工
		if temp.Id == id {
			flag = true
			break
		}
		// 没找到员工
		if temp.Next == nil {
			break
		}
		temp = temp.Next
	}

	// 找到员工就返回
	if flag {
		return temp
	} else {
		fmt.Println("can`t find the employee")
	}
	return
}

// 显示链表的信息
func (el *EmpLink) ShowLink(no int) {
	if el.Head == nil {
		fmt.Printf("链表%d为空\n", no)
		return
	}

	// 遍历当前的链表，并显示数据
	temp := el.Head
	for {
		if temp != nil {
			fmt.Printf("链表%d, 雇员id: %d, 名字: %s", no, temp.Id, temp.Name)
			temp = temp.Next
		} else {
			break
		}

	}
	// 换行处理
	fmt.Println()
}

// 定义hashtable，含有链表数组，数组每个位置存放一条链上的员工信息
type HashTbale struct {
	LinkArr [7]EmpLink
}

// 添加员工的方法
func (ht *HashTbale) Insert(emp *Emp) {
	// 使用散列函数，确定将该员工添加到哪个链表
	linkNo := ht.HashFunc(emp.Id)
	// 使用对应的链表添加
	ht.LinkArr[linkNo].Insert(emp)
}

// 显示hashTable的所有员工
func (ht *HashTbale) ShowAll() {
	for i := 0; i < len(ht.LinkArr); i++ {
		ht.LinkArr[i].ShowLink(i)
	}
}

// 通过员工ID查询
func (ht *HashTbale) FindById(id int) *Emp {
	// 确定该雇员所在的链表
	linkNo := ht.HashFunc(id)
	return ht.LinkArr[linkNo].FindById(id)

}

// 编写一个散列方法
func (ht *HashTbale) HashFunc(id int) int {
	// 得到的值就是对应链表的下标
	return id % 7
}
func main() {

	key := ""
	id := 0
	name := ""

	var hashTable HashTbale

	for {
		fmt.Println("===========================")
		fmt.Println("1. 添加员工")
		fmt.Println("2. 显示员工")
		fmt.Println("3. 查找员工")
		fmt.Println("4. 退出系统")
		fmt.Println("请输入你的选择")
		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("请输入员工id: ")
			fmt.Scanln(&id)
			fmt.Println("请输入员工名字: ")
			fmt.Scanln(&name)
			emp := &Emp{
				Id:   id,
				Name: name,
			}

			hashTable.Insert(emp)
		case "2":
			hashTable.ShowAll()
		case "3":
			fmt.Println("请输入要查找的ID号：")
			fmt.Scanln(&id)
			emp := hashTable.FindById(id)
			if emp != nil {
				emp.ShowMe()
			}

		case "4":
			os.Exit(0)
		}
	}
}

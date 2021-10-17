package main

import "fmt"

type CatNode struct {
	Id   int
	Name string
	Next *CatNode
}

func Insert(head *CatNode, newNode *CatNode) {

	// 第一次添加
	if head.Next == nil {
		head.Id = newNode.Id
		head.Name = newNode.Name
		head.Next = head
		return
	}

	temp := head
	// 找到最后一个节点
	for {
		if temp.Next == head {
			break
		}
		temp = temp.Next
	}

	// 加入到链表中
	temp.Next = newNode
	newNode.Next = head
}

// 遍历环形链表
func List(head *CatNode) {
	temp := head
	if temp.Next == nil {
		fmt.Println("the linked is empty")
		return
	}

	for {
		fmt.Println("猫的名称：", temp.Name)
		if temp.Next == head {
			break
		}
		temp = temp.Next
	}
}

// 删除
func Delete(head *CatNode, id int) *CatNode {
	temp := head
	helper := head

	if temp.Next == nil {
		fmt.Println("the linked is empty")
		return head
	}

	// 如果链表只有一个节点
	if temp.Next == head {
		if temp.Id == id {
			temp.Next = nil
		}
		return head
	}
	// 把helper放到链表的最后
	for {
		if helper.Next == head {
			break
		}
		helper = helper.Next
	}

	// 删除指定的节点
	flag := true
	for {
		if temp.Next == head { // 最后一个节点，还没找到
			break
		}
		if temp.Id == id { // 找到了指定的节点，进行删除
			if temp == head { // 删除的是头节点
				head = head.Next
			}
			helper.Next = temp.Next
			fmt.Println("被删除的小猫：", id)
			flag = false
			break
		}
		temp = temp.Next
		helper = helper.Next
	}

	if flag {
		if temp.Id == id {
			helper.Next = temp.Next
			fmt.Println("被删除的小猫：", id)
		} else {
			fmt.Println("没有这个猫")
		}
	}

	return head
}
func main() {

	// 测试
	head := &CatNode{}

	// 创建猫
	cat1 := &CatNode{
		Id:   1,
		Name: "tom",
	}

	cat2 := &CatNode{
		Id:   2,
		Name: "tom2",
	}

	cat3 := &CatNode{
		Id:   3,
		Name: "tom3",
	}

	Insert(head, cat1)
	Insert(head, cat2)
	Insert(head, cat3)
	List(head)
	head = Delete(head, 1)
	head = Delete(head, 2)
	List(head)

}

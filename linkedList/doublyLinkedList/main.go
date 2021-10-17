package main

import "fmt"

type HeroNode struct {
	Id       int
	Name     string
	NickName string
	Prev     *HeroNode
	Next     *HeroNode
}

// 插入
func Insert(head *HeroNode, newNode *HeroNode) {
	temp := head
	for {
		if temp.Next == nil {
			break
		}
		temp = temp.Next
	}
	temp.Next = newNode
	newNode.Prev = temp
}

// 按照id的大小进行顺序插入
func InsertByOrder(head *HeroNode, newNode *HeroNode) {
	temp := head
	flag := true
	for {
		if temp.Next == nil {
			break
		} else if temp.Next.Id > newNode.Id {
			break
		} else if temp.Next.Id == newNode.Id {
			flag = false
			break
		}
		temp = temp.Next
	}

	if flag {
		newNode.Next = temp.Next
		newNode.Prev = temp
		if temp.Next != nil {
			temp.Next.Prev = newNode
		}
		temp.Next = newNode
	} else {
		fmt.Println("链表已存在该节点")
		return
	}
}

// 删除
func Delete(head *HeroNode, id int) {
	temp := head
	for {
		if temp.Next.Id == id {
			break
		}
		temp = temp.Next
	}

	temp.Next = temp.Next.Next
	// 当要删除的的节点不是最后一个时进行下面的处理
	if temp.Next != nil {
		temp.Next.Prev = temp
	}
}

// 显示所有的节点
func List(head *HeroNode) {
	temp := head
	for {
		if temp.Next == nil {
			break
		}
		temp = temp.Next
		fmt.Printf("[%d, %s, %s]\n", temp.Id, temp.Name, temp.NickName)
	}
}
func main() {

	head := &HeroNode{}

	hero1 := &HeroNode{
		Id:       1,
		Name:     "宋江",
		NickName: "及时雨",
	}

	hero2 := &HeroNode{
		Id:       2,
		Name:     "卢俊义",
		NickName: "玉麒麟",
	}

	hero3 := &HeroNode{
		Id:       3,
		Name:     "林冲",
		NickName: "豹子头",
	}

	// Insert(head, hero1)
	// Insert(head, hero3)
	// Insert(head, hero2)

	InsertByOrder(head, hero1)
	InsertByOrder(head, hero3)
	InsertByOrder(head, hero2)
	List(head)
	Delete(head, 1)
	List(head)
}

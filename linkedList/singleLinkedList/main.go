package main

import "fmt"

type HeroNode struct {
	Id       int
	Name     string
	NickName string
	Next     *HeroNode
}

// 单链表的插入
func Insert(head *HeroNode, newNode *HeroNode) {
	temp := head
	for {
		if temp.Next == nil {
			break
		}
		temp = temp.Next
	}

	temp.Next = newNode
}

// 按顺序插入
func InsertByOrder(head *HeroNode, newNode *HeroNode) {
	temp := head
	flag := true
	for {
		// 已到链表的最后
		if temp.Next == nil {
			break
		} else if temp.Next.Id > newNode.Id { // 已经找到第一个比当前ID大的节点
			break
		} else if temp.Next.Id == newNode.Id { // 已存在相同ID的节点
			flag = false
			break
		}
		temp = temp.Next
	}

	if flag {
		newNode.Next = temp.Next
		temp.Next = newNode
	} else {
		fmt.Println("已存在相同的节点")
		return
	}
}

// 单链表的删除
func Delete(head *HeroNode, id int) {
	temp := head
	for {
		if temp.Next.Id == id {
			break
		}
		temp = temp.Next
	}
	temp.Next = temp.Next.Next
}

// 显示所有的节点信息
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

	InsertByOrder(head, hero1)
	InsertByOrder(head, hero3)
	InsertByOrder(head, hero2)
	List(head)
	Delete(head, 1)
	List(head)
}

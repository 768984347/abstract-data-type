package singly

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	list := NewLinkedList()
	list.Insert(&ListNode{Number: 1}, 0)
	list.Insert(&ListNode{Number: 3}, 1)
	list.Insert(&ListNode{Number: 2}, 1)
	list.Insert(&ListNode{Number: 4}, 3)
	list.Delete(3)
	list.Insert(&ListNode{Number: 4}, 3)
	list.Delete(0)
	list.Insert(&ListNode{Number: 0}, 0)
	list.Insert(&ListNode{Number: 1}, 1)
	list.Insert(&ListNode{Number: 5}, list.Count())
	var num *ListNode
	num = list.GetNode(0)
	if num.Number != 0 {
		t.Errorf("GetNode = %d,want %d", num.Number, 0)
	}
	num = list.GetNode(4)
	if num.Number != 4 {
		t.Errorf("GetNode = %d,want %d", num.Number, 4)
	}
	num = list.GetNode(5)
	if num.Number != 5 {
		t.Errorf("GetNode = %d,want %d", num.Number, 5)
	}
	num = list.GetNodeByNum(5)
	if num.Number != 5 {
		t.Errorf("GetNode = %d,want %d", num.Number, 5)
	}
	index := list.GetNodeIndexByNum(100)
	if index != -1 {
		t.Errorf("GetNode index fail: 100")
	}
	index = list.GetNodeIndexByNum(5)
	if index != 5 {
		t.Errorf("GetNode index fail: 5")
	}
	Display(list)
}

/**
展示链表数据
*/
func Display(list *LinkedList) {
	p := list.Head
	for p != nil {
		fmt.Println(p.Number)
		p = p.Next
	}
	if list.Head != nil {
		fmt.Printf("头结点: %d\n", list.Head.Number)
	} else {
		fmt.Printf("头结点: null\n")
	}
	if list.Tail != nil {
		fmt.Printf("尾结点: %d\n", list.Tail.Number)
	} else {
		fmt.Printf("尾结点: null\n")
	}
	fmt.Printf("Count: %d\n", list.Count())
}

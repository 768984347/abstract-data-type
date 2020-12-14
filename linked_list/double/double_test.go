package circular

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	double := NewLinkedList()
	fmt.Println("===添加测试===")
	double.Insert(&ListNode{Number: 20}, 0)
	double.Insert(&ListNode{Number: 30}, 1)
	double.Insert(&ListNode{Number: 50}, 2)
	display(double)
	fmt.Println("===插入头结点测试===")
	double.Insert(&ListNode{Number: 10}, 0)
	display(double)
	fmt.Println("===插入中间结点测试===")
	double.Insert(&ListNode{Number: 40}, 3)
	display(double)
	fmt.Println("===插入尾结点测试===")
	double.Insert(&ListNode{Number: 60}, double.Count())
	display(double)
	fmt.Println("===获取结点测试===")
	num10 := double.GetNode(0)
	if num10.Number != 10 {
		t.Errorf("GetNode = %d, want %d", num10.Number, 10)
	}
	num60 := double.GetNode(double.Count() - 1)
	if num60.Number != 60 {
		t.Errorf("GetNode = %d, want %d", num60.Number, 60)
	}
	num30 := double.GetNode(2)
	if num30.Number != 30 {
		t.Errorf("GetNode = %d, want %d", num30.Number, 30)
	}
	fmt.Println("===根据结点插入测试,作为前驱结点插入===")
	double.InsertByNode(&ListNode{Number: 0}, double.GetNode(0), 2)
	double.InsertByNode(&ListNode{Number: 55}, double.GetNode(double.Count()-1), 2)
	double.InsertByNode(&ListNode{Number: 22}, double.GetNode(3), 2)
	display(double)
	fmt.Println("===根据结点插入测试,作为后继结点插入===")
	double.InsertByNode(&ListNode{Number: 26}, double.GetNode(3), 1)
	double.InsertByNode(&ListNode{Number: 5}, double.GetNode(0), 1)
	double.InsertByNode(&ListNode{Number: 70}, double.GetNode(double.Count()-1), 1)
	display(double)
	fmt.Println("===删除测试===")
	double.Delete(4)
	double.Delete(0)
	double.Delete(double.Count() - 1)
	display(double)
	fmt.Println("===删除结点测试===")
	double.DeleteNode(double.GetNode(3))
	double.DeleteNode(double.GetNode(0))
	double.DeleteNode(double.GetNode(double.Count() - 1))
	display(double)
}

func display(list *LinkedList) {
	p := list.Head
	for i := 0; i < list.Count(); i++ {
		fmt.Printf("下标: %d,", i)
		if p.Prev != nil {
			fmt.Printf("前驱结点: %d， ", p.Prev.Number)
		} else {
			fmt.Printf("前驱结点: null， ")
		}
		fmt.Printf("当前结点: %d，", p.Number)
		if p.Next != nil {
			fmt.Printf("后继结点: %d，", p.Next.Number)
		} else {
			fmt.Printf("后继结点: null，")
		}
		fmt.Printf("\n")
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

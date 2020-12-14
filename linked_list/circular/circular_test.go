package circular

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	circular := NewLinkedList()
	fmt.Println("==正常测试==")
	circular.Insert(&ListNode{Number: 0}, 0)
	circular.Insert(&ListNode{Number: 1}, 1)
	circular.Insert(&ListNode{Number: 3}, 2)
	circular.Insert(&ListNode{Number: 2}, 2)
	display(circular)
	fmt.Println("==新增头结点==")
	circular.Insert(&ListNode{Number: -1}, 0)
	display(circular)
	fmt.Println("==删除头结点==")
	circular.Delete(0)
	display(circular)
	fmt.Println("==删除尾结点==")
	circular.Delete(circular.Count() - 1)
	display(circular)
	if num := circular.GetNode(0).Number; num != 0 {
		t.Errorf("GetNode = %d,want %d", num, 0)
	}
	if num := circular.GetNode(1).Number; num != 1 {
		t.Errorf("GetNode = %d,want %d", num, 1)
	}
	if num := circular.GetNode(2).Number; num != 2 {
		t.Errorf("GetNode = %d,want %d", num, 2)
	}
}

func display(list *LinkedList) {
	p := list.Head
	for i := 0; i < list.Count(); i++ {
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
	if list.Tail.Next != nil {
		fmt.Printf("尾结点next指向: %d\n", list.Tail.Next.Number)
	} else {
		fmt.Printf("尾结点next指向 null\n")
	}
	fmt.Printf("Count: %d\n", list.Count())
}

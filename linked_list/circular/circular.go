package circular

/*
循环链表实现
@language Golang
*/

/**
循环链表结点(元素)
*/
type ListNode struct {
	Number int       //数据
	Next   *ListNode //指向下一个结点的指针
}

/**
循环链表结构体
*/
type LinkedList struct {
	Head *ListNode //头指针,固定指向第一个结点
	Tail *ListNode //尾指针,固定指向最后一个结点
	Size int       //当前链表大小
}

/**
初始化循环链表
*/
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

/**
循环链表插入结点
@param node 要插入的结点
@param index 在index和index-1元素之间插入一个元素
*/
func (list *LinkedList) Insert(node *ListNode, index int) {
	if index > list.Size || index < 0 {
		panic("index out of range")
	}
	if index == 0 { //向表头添加结点
		if list.Head == nil {
			list.Head = node //初始化头指针
			list.Tail = node //初始化尾指针
		} else {
			list.Head, node.Next = node, list.Head
		}
		list.Tail.Next = list.Head //尾结点指向头指针
	} else if index == list.Size { //向表尾添加结点
		list.Tail.Next = node //原尾结点next指向新尾结点
		list.Tail = node      //尾指针指向新的尾结点
		node.Next = list.Head //尾结点next重新指向头结点
	} else { //向中间插入结点
		var prev = list.Head
		for i := 1; i < index; i++ { //找到要插入点的前一个结点
			prev = prev.Next
		}
		/**
		等同于:
			node.Next = prev.Next
			prev.Next = node
		*/
		prev.Next, node.Next = node, prev.Next
	}
	list.Size++
	return
}

/**
循环链表删除结点
@param index 要删除结点的下标
*/
func (list *LinkedList) Delete(index int) {
	if index >= list.Size || index < 0 {
		panic("index out of range")
	}
	if index == 0 { //处理头结点的删除
		list.Head = list.Head.Next
		if list.Head == nil { //当链表被清空时,尾结点也置为null
			list.Tail = nil
		} else {
			list.Tail.Next = list.Head //尾结点重新指向头结点
		}
	} else {
		var prev, curr, next *ListNode
		prev = list.Head
		for i := 1; i < index; i++ { //找到要删除的前一个结点
			prev = prev.Next
		}
		curr = prev.Next      //要删除的目标结点
		if curr.Next != nil { //要删除元素的下一个结点赋值给next
			next = curr.Next
		}
		prev.Next = next
		if next == list.Head { //删除最后一个元素 重新赋值尾结点
			list.Tail = prev
			list.Tail.Next = list.Head
		}
	}
	list.Size--
	return
}

/**
根据下标获取循环链表结点
@param index 要获取结点的下标
*/
func (list *LinkedList) GetNode(index int) *ListNode {
	if index >= list.Size || index < 0 {
		panic("index out of range")
	}
	var p = list.Head
	for i := 0; i < index; i++ {
		p = p.Next
	}
	return p
}

/**
获取循环链表数量
*/
func (list *LinkedList) Count() int {
	return list.Size
}

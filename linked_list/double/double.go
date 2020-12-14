package circular

/*
双向链表实现
@language Golang
*/

/**
双向链表结点(元素)
*/
type ListNode struct {
	Number int       //数据
	Next   *ListNode //指向后继结点的指针
	Prev   *ListNode //指向前驱结点的指针
}

/**
双向链表结构体
*/
type LinkedList struct {
	Head *ListNode //头指针,固定指向第一个结点
	Tail *ListNode //尾指针,固定指向最后一个结点
	Size int       //当前链表大小
}

/**
初始化双向链表
*/
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

/**
双向链表插入结点
@param node 要插入的结点
@param index 在index和index-1元素之间插入一个元素
*/
func (list *LinkedList) Insert(node *ListNode, index int) {
	if index > list.Count() || index < 0 {
		panic("index out of range")
	}
	if index == 0 {
		if list.Head == nil { //初始化头尾指针
			list.Head = node
			list.Tail = node
		} else {
			list.Head.Prev = node //原头结点prev指向node
			node.Next = list.Head //node成为新的头结点 next指向原头结点
			list.Head = node      //头指针重新指向新的头结点
		}
	} else if index == list.Count() { //向尾部添加结点
		node.Prev = list.Tail //node成为新的尾结点 prev指向原尾结点
		list.Tail.Next = node //原尾结点next指向node
		list.Tail = node      //尾指针重新指向新的尾结点
	} else {
		var prev, curr *ListNode
		prev = list.Head
		for i := 1; i < index; i++ { //找要添加结点的上一个结点
			prev = prev.Next
		}
		curr = prev.Next //要插入位置的结点
		curr.Prev = node //node成为curr的前驱结点
		node.Prev = prev //prev成为node的前驱结点
		node.Next = curr //curr成为node的后继结点
		prev.Next = node //node成为prev的后继结点
	}
	list.Size++
	return
}

/**
通过targetNode添加新的结点
@param addNode 要添加的结点
@param targetNode 目标结点
@param addType 1 => addNode作为targetNode的后继结点 2 =>  addNode作为targetNode的前驱结点
*/
func (list *LinkedList) InsertByNode(addNode *ListNode, targetNode *ListNode, addType int) {
	if addType == 1 {
		if targetNode.Next == nil {
			list.Tail = addNode //targetNode是尾结点,重置尾指针
		} else {
			targetNode.Next.Prev = addNode
		}
		addNode.Next = targetNode.Next
		addNode.Prev = targetNode
		targetNode.Next = addNode
	} else {
		if targetNode.Prev == nil {
			list.Head = addNode //targetNode是头结点,重置尾指针
		} else {
			targetNode.Prev.Next = addNode
		}
		addNode.Next = targetNode
		addNode.Prev = targetNode.Prev
		targetNode.Prev = addNode
	}
	list.Size++
}

/**
双向链表删除结点
@param index 要删除结点的下标
*/
func (list *LinkedList) Delete(index int) {
	if index >= list.Count() || index < 0 {
		panic("index out of range")
	}
	var prev, curr, next *ListNode
	curr = list.Head
	for i := 0; i < index; i++ { //找到要删除的结点
		curr = curr.Next
	}
	prev = curr.Prev
	next = curr.Next
	if prev != nil { //判断不是头结点
		prev.Next = next //将前驱结点的后继改为next
	} else {
		list.Head = next //修改头结点指向next
	}
	if next != nil { //判断不是尾结点
		next.Prev = prev //将后继结点的前驱改为prev
	} else {
		list.Tail = prev //修改尾结点指向prev
	}
	list.Size--
	return
}

/**
删除指定结点
@param node 要删除的结点 node必须是list中的结点
*/
func (list *LinkedList) DeleteNode(node *ListNode) {
	var prev, next *ListNode
	prev = node.Prev
	next = node.Next
	if prev == nil {
		list.Head = next //删除头结点重置头指针
	} else {
		prev.Next = next
	}
	if next == nil {
		list.Tail = prev //删除尾结点重置尾指针
	} else {
		next.Prev = prev
	}
	list.Size--
}

/**
根据下标获取链表结点
@param index 要获取结点的下标
*/
func (list *LinkedList) GetNode(index int) *ListNode {
	if index < 0 || index >= list.Count() {
		panic("index out of range")
	}
	p := list.Head
	for i := 0; i < index; i++ { //找到查找的结点
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

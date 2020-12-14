package singly

import (
	"abstract-data-type/linked_list/singly"
	"fmt"
	"testing"
)

type TestItem struct {
	number int
	val    string
}

func Test(t *testing.T) {
	//初始化5个结点
	arr := [5]TestItem{
		{number: 1, val: "one"},
		{number: 2, val: "two"},
		{number: 3, val: "three"},
		{number: 4, val: "four"},
		{number: 5, val: "five"},
	}
	fmt.Println("初始化数据")
	storage := NewStorage(5) //设置最大保存5个数据
	for _, v := range arr {
		storage.Set(v.number, v.val) //将结点存入到storage,此时顺序为5,4,3,2,1
	}
	displayLinkedList(storage.linkedList)
	fmt.Println("当前map数据")
	fmt.Println(storage.data)
	fmt.Println("查询2的value,2提到表首")
	val, err := storage.Get(2)
	if err != nil {
		panic(err)
	}
	if val != "two" {
		t.Errorf("Get Val = %s,need %s", val, "two")
	}
	displayLinkedList(storage.linkedList)
	fmt.Println("查询5的value,5提到表首")
	val, err = storage.Get(5)
	if err != nil {
		panic(err)
	}
	if val != "five" {
		t.Errorf("Get Val = %s,need %s", val, "five")
	}
	displayLinkedList(storage.linkedList)
	fmt.Println("设置6,6不存在新增,新增之后最大保留数超过5个,1最少访问被删除,6添加到表首")
	storage.Set(6, "six")
	displayLinkedList(storage.linkedList)
	fmt.Println("当前map数据")
	fmt.Println(storage.data)
	fmt.Println("设置7,7不存在新增,新增之后最大保留数超过5个,3最少访问被删除,7添加到表首")
	storage.Set(7, "seven")
	displayLinkedList(storage.linkedList)
	fmt.Println("当前map数据")
	fmt.Println(storage.data)
	fmt.Println("设置4,4已经存在,4添加到表首")
	storage.Set(4, "four")
	displayLinkedList(storage.linkedList)
	fmt.Println("当前map数据")
	fmt.Println(storage.data)
}

func displayLinkedList(list *singly.LinkedList) {
	p := list.Head
	index := 0
	for p != nil {
		fmt.Printf("index=%d,num=%d\n", index, p.Number)
		p = p.Next
		index++
	}
	fmt.Printf("count: %d\n", list.Count())
}

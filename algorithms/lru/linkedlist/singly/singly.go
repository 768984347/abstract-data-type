package singly

import (
	"abstract-data-type/linked_list/singly"
	"errors"
)

/**
存储结构体
单链表实现lru算法
*/
type Storage struct {
	linkedList *singly.LinkedList //单链表用于实现lru淘汰算法
	data       map[int]string     //真实的数据保存
	MaxSize    int                //最大数据数量
}

func NewStorage(maxSize int) *Storage {
	return &Storage{linkedList: &singly.LinkedList{}, MaxSize: maxSize, data: map[int]string{}}
}

/**
获取数据
*/
func (s *Storage) Get(number int) (string, error) {
	index := s.linkedList.GetNodeIndexByNum(number) //根据数据获取下标
	if index == -1 {
		return "", errors.New("number not found")
	}
	//数据被访问就重新添加到链表头
	node := s.linkedList.GetNode(index)
	s.linkedList.Delete(index)
	s.linkedList.Insert(node, 0)
	return s.data[number], nil
}

/**
设置数据
*/
func (s *Storage) Set(number int, value string) {
	index := s.linkedList.GetNodeIndexByNum(number)
	if index == -1 {
		if s.linkedList.Count()+1 > s.MaxSize { //如果单链表内没有number,需要清理出空间
			s.clear(1)
		}
	} else {
		s.linkedList.Delete(index) //活跃数据先删除，之后重新放置在表头
	}
	s.linkedList.Insert(&singly.ListNode{Number: number}, 0) //新设置的数据在单链表头
	s.data[number] = value
}

/**
删除数据
*/
func (s *Storage) Delete(number int) {
	index := s.linkedList.GetNodeIndexByNum(number)
	if index == -1 {
		panic("number not found")
	}
	s.linkedList.Delete(index)
	delete(s.data, number)
}

/**
清除指定数量的数据
*/
func (s *Storage) clear(clearNum int) {
	if clearNum < 0 {
		panic("clearNum can not less than zero")
	}
	linkedSize := s.linkedList.Count()
	if clearNum > linkedSize {
		clearNum = linkedSize //如果要清除所有数据,clearNum为单链表的数量
	}
	linkedLastIndex := linkedSize - 1
	for i := 0; i < clearNum; i++ {
		node := s.linkedList.GetNode(linkedLastIndex) //获取结点
		s.linkedList.Delete(linkedLastIndex)          //总是删除单链表尾部的数据
		delete(s.data, node.Number)                   //删除data内的数据
		linkedLastIndex--
	}
}

package stack

import (
	"errors"
)

/**
栈
@特性
先进后出
@method Init 初始化栈
@method Push 入栈
@method Pop 出栈
@method IsEmpty 判断栈是否为空
@method IsFull 判断栈是否已满
@method GetLength 获取栈长度
@method Top 获取栈首个元素
@author pxb
*/

type Stack struct {
	data []interface{}
	len  int
	size int
}

func Init(maxLen int) *Stack {
	return &Stack{make([]interface{}, maxLen), 0, maxLen}
}

func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("empty stack")
	}
	defer func() {
		s.len--
	}()
	return s.data[s.len-1], nil
}

func (s *Stack) Push(item interface{}) bool {
	if s.IsFull() {
		return false
	}
	s.data[s.len] = item
	s.len++
	return true
}

func (s *Stack) IsEmpty() bool {
	if s.len > 0 {
		return false
	}
	return true
}

func (s *Stack) IsFull() bool {
	if s.len >= s.size {
		return true
	}
	return false
}

func (s *Stack) GetLength() int {
	return s.len
}

func (s *Stack) Top() interface{} {
	if s.IsEmpty() {
		panic("empty stack")
	}
	return s.data[s.len-1]
}

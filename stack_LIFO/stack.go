package stack

import (
	"errors"
)

/**
栈
@特性
先进后出
@实现原理
只需要保留当前数组的len作为下一个插入的元素index即可
@method Init 初始化栈
@method Push 入栈
@method Pop 出栈
@method IsEmpty 判断队列是否为空
@method IsFull 判断队列是否已满
@method GetLength 获取队列长度
*/

type Stack struct {
	data []int
	len int
}

func Init(maxLen int) *Stack {
	return &Stack{make([]int, maxLen), 0}
}

func (this *Stack) Pop() (int, error) {
	if this.IsEmpty() {
		return -1, errors.New("stack is empty")
	}
	defer func() {
		this.len--
	}()
	return this.data[this.len - 1], nil
}

func (this *Stack) Push(number int) bool {
	if this.IsFull() {
		return false
	}
	this.data[this.len] = number
	this.len++
	return true
}

func (this *Stack) IsEmpty() bool {
	if this.len > 0 {
		return false
	}
	return true
}

func (this *Stack) IsFull() bool {
	if this.len >= len(this.data) {
		return true
	}
	return false
}

func (this *Stack) GetLength() int {
	return this.len
}
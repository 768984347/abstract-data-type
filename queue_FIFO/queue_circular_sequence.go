package queue

import "fmt"

/**
循环队列
@特性
先进先出
@实现原理
通过记录length和firstIndex(队列的首个元素index)来计算可以插入的位置,可以做到在一个固定长度的数组中实现循环队列,
相比于简单队列而言能够节约内存
@method Constructor 初始化队列
@method EnQueue 入队
@method DeQueue 出队
@method Front 获取第一个队列元素
@method Rear 获取最后一个队列元素
@method Rear 获取最后一个队列元素
@method IsEmpty 判断队列是否为空
@method IsFull 判断队列是否已满
@method getRealIndex 如果index超过length返回正确的index(超过部分再从头开始计算)
@method GetLength 获取队列长度
@method ShowList 调试方法
 */

type MyCircularQueue struct {
	data []int
	firstIndex int
	length int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{make([]int, k), 0, 0}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.length >= len(this.data) {
		return false
	}

	defer func() {
		this.length++
	}()

	nextIndex := this.getRealIndex(this.firstIndex + this.length)
	this.data[nextIndex] = value

	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	defer func() { //将长度-1
		this.length--
		if this.IsEmpty() {
			this.firstIndex = 0
		} else {
			this.firstIndex = this.getRealIndex(this.firstIndex + 1)
		}
	}()

	this.data[this.firstIndex] = 0

	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.firstIndex]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	index := this.getRealIndex(this.length + this.firstIndex - 1)
	return this.data[index]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	if this.length > 0 {
		return false
	}
	return true
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	if this.length >= len(this.data) {
		return true
	}
	return false
}

func (this *MyCircularQueue) getRealIndex(index int) int {
	if index >= len(this.data) {
		return index - len(this.data)
	}
	return index
}

func (this *MyCircularQueue) ShowList() {
	fmt.Printf("%v", this.data)
}

func (this *MyCircularQueue) GetLength() int {
	return this.length
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
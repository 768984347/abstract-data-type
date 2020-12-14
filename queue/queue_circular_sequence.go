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
@author pxb
*/

type MyCircularQueue struct {
	data       []int
	firstIndex int
	length     int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{make([]int, k), 0, 0}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (queue *MyCircularQueue) EnQueue(value int) bool {
	if queue.length >= len(queue.data) {
		return false
	}

	defer func() {
		queue.length++
	}()

	nextIndex := queue.getRealIndex(queue.firstIndex + queue.length)
	queue.data[nextIndex] = value

	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (queue *MyCircularQueue) DeQueue() bool {
	if queue.IsEmpty() {
		return false
	}
	defer func() { //将长度-1
		queue.length--
		if queue.IsEmpty() {
			queue.firstIndex = 0
		} else {
			queue.firstIndex = queue.getRealIndex(queue.firstIndex + 1)
		}
	}()

	queue.data[queue.firstIndex] = 0

	return true
}

/** Get the front item from the queue. */
func (queue *MyCircularQueue) Front() int {
	if queue.IsEmpty() {
		return -1
	}
	return queue.data[queue.firstIndex]
}

/** Get the last item from the queue. */
func (queue *MyCircularQueue) Rear() int {
	if queue.IsEmpty() {
		return -1
	}
	index := queue.getRealIndex(queue.length + queue.firstIndex - 1)
	return queue.data[index]
}

/** Checks whether the circular queue is empty or not. */
func (queue *MyCircularQueue) IsEmpty() bool {
	if queue.length > 0 {
		return false
	}
	return true
}

/** Checks whether the circular queue is full or not. */
func (queue *MyCircularQueue) IsFull() bool {
	if queue.length >= len(queue.data) {
		return true
	}
	return false
}

func (queue *MyCircularQueue) getRealIndex(index int) int {
	if index >= len(queue.data) {
		return index - len(queue.data)
	}
	return index
}

func (queue *MyCircularQueue) ShowList() {
	fmt.Printf("%v", queue.data)
}

func (queue *MyCircularQueue) GetLength() int {
	return queue.length
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

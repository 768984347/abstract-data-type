package queue

import "errors"

/**
队列
顺序存储结构实现队列
@特性
先进先出
@method EnQueue 进列
@method DeQueue 出队
@method IsDataExists 检查数据是否在队列中存在
@method GetFirst 获取位于队列首的数据
@method GetLength 获取队列当前的元素数量
@author pxb
*/

type Queue struct {
	data     []interface{}
	head     int
	tail     int
	capacity int
	count    int
}

/**
初始化队列
*/
func Init(capacity int) *Queue {
	return &Queue{
		data:     make([]interface{}, capacity),
		capacity: capacity,
	}
}

/**
入队列
*/
func (queue *Queue) EnQueue(data interface{}) bool {
	if queue.count == queue.capacity { //队列满
		return false
	}
	if queue.tail == queue.capacity { //到尾部，重新指向头部
		queue.tail = 0
	}
	queue.data[queue.tail] = data
	queue.tail++
	queue.count++
	return true
}

/**
出队列
*/
func (queue *Queue) DeQueue() (interface{}, error) {
	if queue.count == 0 {
		return nil, errors.New("empty queue")
	}
	if queue.head == queue.capacity { //到尾部，重新指向头部
		queue.head = 0
	}
	res := queue.data[queue.head]
	queue.head++
	queue.count--
	return res, nil
}

/**
检查数据是否存在
*/
func (queue *Queue) IsDataExists(data interface{}) bool {
	for i := queue.head; i < queue.tail; i++ {
		if queue.data[i] == data {
			return true
		}
	}
	return false
}

/**
获取首个元素
*/
func (queue *Queue) GetFirst() (interface{}, error) {
	if queue.count > 0 {
		return queue.data[queue.head], nil
	}
	return nil, errors.New("empty queue")
}

/**
获取队列存在元素的数量
*/
func (queue *Queue) GetLength() int {
	return queue.count
}

package queue

/**
简单队列
@实现原理
在出队列的时候将firstIndex自增来记录当前队列的首个元素位置,不进行元素删除的操作
在统计队列元素个数的时候(总的元素数量 - firstIndex)可以得到当前队列的元素总数
@method EnQueue 进列
@method DeQueue 出队
@method IsDataExists 检查数据是否在队列中存在
@method GetFirst 获取位于队列首的数据
@method GetLength 获取队列当前的元素数量
 */

import "errors"

type Queue struct {
	data []interface{}
	firstIndex int
}

func (queue *Queue) EnQueue(data interface{}) bool {
	queue.data = append(queue.data, data)
	return true
}

func (queue *Queue) DeQueue() (interface{}, error) {
	if queue.GetLength() <= 0 {
		return nil, errors.New("empty queue")
	}

	defer func() {
		queue.firstIndex++
	}()

	return queue.data[queue.firstIndex], nil
}

func (queue *Queue) IsDataExists(data interface{}) bool {
	for _, queueData := range queue.data {
		if queueData == data {
			return true
		}
	}
	return false
}

func (queue *Queue) GetFirst() (interface{}, error) {
	if queue.GetLength() > 0 {
		return queue.data[queue.firstIndex], nil
	}
	return nil, errors.New("empty queue")
}

func (queue *Queue) GetLength() int {
	return len(queue.data) - queue.firstIndex
}
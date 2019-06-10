package examination

import (
	"abstract_data_type/queue_FIFO"
)

type MovingAverage struct {
	queue queue.MyCircularQueue
	sum   int
}

func Init(maxLength int) *MovingAverage {
	if maxLength <= 0 {
		panic("maxLength less than 0")
	}
	return &MovingAverage{queue.Constructor(maxLength), 0}
}

func (this *MovingAverage) Next(number int) {
	if this.queue.IsFull() {
		front := this.queue.Front()
		if !this.queue.DeQueue() {
			panic("function DeQueue error")
		}
		this.sum -= front
	}

	if !this.queue.EnQueue(number) {
		panic("function EnQueue error")
	}
	this.sum += number
}

func (this *MovingAverage) GetAverage() float64 {
	if this.sum == 0 {
		return 0
	}
	return float64(this.sum) / float64(this.queue.GetLength())
}
package main

import (
	queue2 "abstract_data_type/queue_FIFO"
	"fmt"
)

func main() {
	queue := queue2.Queue{}

	queue.EnQueue(123)
	queue.EnQueue("hello")

	first, err := queue.GetFirst()

	if err != nil {
		panic(err)
	}

	fmt.Printf("len = %d\n first = %v\n", queue.GetLength(), first)

	fmt.Printf("is 123 exists ? %v\n", queue.IsDataExists(123))
	fmt.Printf("is world exists ? %v\n", queue.IsDataExists("world"))

	for queue.GetLength() > 0 {
		data, err := queue.DeQueue()
		if err != nil {
			panic(err)
		}
		fmt.Printf("item = %v\n", data)
	}

	fmt.Printf("now len = %d\n", queue.GetLength())

	fmt.Println("==============================")

	circularQueue := queue2.Constructor(6)

	circularQueue.EnQueue(6)
	circularQueue.Rear()
	circularQueue.Rear()
	circularQueue.DeQueue()
	circularQueue.EnQueue(5)
	circularQueue.Rear()
	circularQueue.DeQueue()
	fmt.Println(	circularQueue.Front())
	circularQueue.DeQueue()
	circularQueue.DeQueue()
	circularQueue.DeQueue()
}
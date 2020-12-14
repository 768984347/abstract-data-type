package queue

import (
	"fmt"
	"testing"
)

func TestEnQueue(t *testing.T) {
	arr := [8]interface{}{"a", "b", "c", "d", "e", "f", "g"}
	queue := Init(4)
	display(queue)
	queue.EnQueue(arr[0])
	queue.EnQueue(arr[1])
	queue.EnQueue(arr[2])
	queue.EnQueue(arr[3])
	suc := queue.EnQueue(arr[4])
	if suc != false {
		t.Errorf("Range Out Of queue. %s", arr[4])
	}
	display(queue)
}

func TestDeQueue(t *testing.T) {
	arr := [8]interface{}{"a", "b", "c", "d", "e", "f", "g"}
	queue := Init(4)
	for i := 0; i < 4; i++ {
		queue.EnQueue(arr[i])
	}
	for i := 0; i < 4; i++ {
		res, err := queue.DeQueue()
		if err != nil {
			t.Error(err.Error())
		}
		if res != arr[i] {
			t.Errorf("Want %s,Get %s", arr[i], res)
		}
	}
	res, err := queue.DeQueue()
	if res != nil || err.Error() != "empty queue" {
		t.Errorf("Out Range Of Head")
	}
	display(queue)
}

func TestFrontier(t *testing.T) {
	arr := [8]interface{}{"a", "b", "c", "d", "e", "f", "g"}
	queue := Init(4)
	for i := 0; i < 4; i++ {
		queue.EnQueue(arr[i])
	}
	for i := 0; i < 4; i++ {
		queue.DeQueue()
	}
	queue.EnQueue(arr[5])
	res, err := queue.DeQueue()
	if err != nil {
		t.Error(err.Error())
	}
	if res != arr[5] {
		t.Errorf("Want %d,Get %s", arr[5], res)
	}
	display(queue)
}

func TestCross(t *testing.T) {
	arr := [8]interface{}{"a", "b", "c", "d", "e", "f", "g"}
	queue := Init(4)
	for i := 0; i < len(arr)-1; i++ {
		suc := queue.EnQueue(arr[i])
		if suc != true {
			t.Errorf("EnQueue Fail: %s", arr[i])
		}
		res, err := queue.DeQueue()
		if err != nil {
			t.Error(err.Error())
		}
		if res != arr[i] {
			t.Errorf("Want %d,Get %s", arr[i], res)
		}
	}
	display(queue)
}

func display(queue *Queue) {
	fmt.Println(queue.data)
	fmt.Printf("head: %d, tail: %d,count: %d, cap: %d\n", queue.head, queue.tail, queue.count, queue.capacity)
}

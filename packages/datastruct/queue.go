package datastruct

import "fmt"

// Queue (FIFO) data structure used in breadth-first traversal of a tree data structure
type Queue struct {
	front *QueueElement
	rear  *QueueElement
	size  int
}

// QueueElement is defining the structure of each element in the queue
type QueueElement struct {
	value interface{}
	next  *QueueElement
}

// Enqueue a new element at the end of the queueu
func (queue *Queue) Enqueue(data interface{}) {
	el := &QueueElement{value: data}
	if queue.size == 0 {
		queue.front = el
		queue.rear = queue.front
		queue.size = 1
		return
	}

	queue.rear.next = el
	queue.rear = el
	queue.size++
}

// Dequeue the element ot the front of the queueu
func (queue *Queue) Dequeue() interface{} {
	if queue.size == 0 {
		return nil
	}

	front := queue.front
	queue.front = queue.front.next
	queue.size--

	if queue.size == 0 {
		queue.front = nil
		queue.rear = nil
	}

	return front.value
}

// Len returns the length of the queue
func (queue *Queue) Len() int {
	return queue.size
}

// Implementation of String() method from Stringer interface
func (queue *Queue) String() string {
	el := queue.front

	var stringBuilder string
	stringBuilder = fmt.Sprintln("BEGIN: Queue elements")
	for el != nil {
		stringBuilder = fmt.Sprintln(stringBuilder, el.value)
		el = el.next
	}
	stringBuilder = fmt.Sprintln(stringBuilder, "END")
	return stringBuilder
}

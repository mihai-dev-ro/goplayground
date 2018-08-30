package datastruct

import "fmt"

// Stack (LIFO) data structure used in depth-first traversal of a tree data structure
type Stack struct {
	top  *StackElement
	size int
}

// StackElement defines the structure of each element in the stack
type StackElement struct {
	next  *StackElement
	value interface{}
}

// Push new element on the stack
func (stack *Stack) Push(data interface{}) {
	el := &StackElement{value: data, next: stack.top}
	stack.top = el
	stack.size++
}

// Pop the element at the top of the stack
func (stack *Stack) Pop() interface{} {
	if stack.size == 0 {
		return nil
	}

	result := stack.top
	stack.top = result.next
	stack.size--
	return result.value
}

// Peek allows to see the element at the top of the stack without removing it
func (stack *Stack) Peek() interface{} {
	if stack.size == 0 {
		return nil
	}

	return stack.top.value
}

// Len returns the number of elements in the stack
func (stack *Stack) Len() int {
	return stack.size
}

// Implementation of String() method from Stringer interface
func (stack *Stack) String() string {
	el := stack.top

	var stringBuilder string
	stringBuilder = fmt.Sprintln("BEGIN: Stack elements")
	for el != nil {
		stringBuilder = fmt.Sprintln(stringBuilder, el.value)
		el = el.next
	}
	stringBuilder = fmt.Sprintln(stringBuilder, "END")
	return stringBuilder
}

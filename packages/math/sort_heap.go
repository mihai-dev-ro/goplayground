package math

import (
	 "fmt"
)

// Implementation of Heap Sort algorithm
// Heap Sort is an algorithm that implies building a heap data structure (which is a tree that 
// has the nodes arranged in a certain/chosen order: max-heap or min-heap. A max-heap tree has the property
// that each parent node has a greater value that each of child nodes) and then progressively moving the 
// root of the tree to append/prepend the sorted list. 
// After each move of the root of the tree to the sorted list, we must re-arrange the heap data structure so 
// that it remains a valid heap
//
// Complexity analysis
//    - time complexity: 
//    - space complexity:
func SortHeap(list []int) []int {
	//fmt.Printf("Input list: %v\n", list)

	// empty or 1-item list; just return it as there is nothing to sort
	if len(list) < 2 {
		return list
	}

	// create the max-heap data structure
	tree := heapifyList(list)

	// extract values for the heap tree into the sorted list
	result := make([]int, len(list))
	index := len(result) - 1 
	for tree != nil && index >= 0 {
		result[index] = tree.value
		index--
		tree = tree.removeRootAndRearrange2()
	}

	//fmt.Printf("Heapsort sorted list: %v\n", result)
	return result
}

func heapifyList(list []int) *HeapTreeNode {

	var treeRoot *HeapTreeNode

	var availableParentsQueue = new(Queue)
	var availableParent *HeapTreeNode


	for _, val := range list {
		if treeRoot == nil {
			treeRoot = &HeapTreeNode{ value:val }
			availableParentsQueue.enqueue(&QueueElement{ value:treeRoot })
			continue
		} else {
			if availableParent == nil || availableParent.numberOfChildren() == 2 {
				availableParent = availableParentsQueue.dequeue().value
			}
			newNode := insertValue(availableParent, val)
			// new nodes added to the queue in the order they are added
			availableParentsQueue.enqueue(&QueueElement{ value:newNode })
		}
	}

	return treeRoot
}

func insertValue(parentNode *HeapTreeNode, val int) *HeapTreeNode {
	
	if parentNode == nil {
		panic("Parent node should never be empty. Is the liaison required by the new node to be created")
	}

	// add the new value to the data structure
	var insertedNode = parentNode.addChild(val)
	

	// 2. re-arrange nodes so that the data structure remains a valid max-heap data structure
	isSwapPerformed := true
	node := insertedNode
	for isSwapPerformed && node.parent != nil {
		
		if node.value > node.parent.value {
			node.swap(node.parent)
			isSwapPerformed = true
			node = node.parent			
		} else {
			isSwapPerformed = false
		}
	}

	return insertedNode
}

// Implementation of Queue (FIFO) data structure used in breadth-first traversal of the Heap data 
type Queue struct {
	front *QueueElement
	rear *QueueElement
	size int
} 

// Enqueue a new element at the end of the queueu
func (this *Queue) enqueue(n *QueueElement) {
	if this.size == 0 {
		this.front = n
		this.rear = this.front
		this.size = 1
		return
	}

	this.rear.next = n
	this.rear = n
	this.size++
}

// Dequeue the element ot the front of the queueu 
func (this *Queue) dequeue() *QueueElement {
	if this.size == 0 {
		return nil
	}

	front := this.front
	this.front = this.front.next
	this.size--

	if this.size == 0 {
		this.front = nil
		this.rear = nil
	}

	return front
}

// Get the length of the queue
func (this *Queue) len() int {
	return this.size
}

// implement method of stringer interface
func (this *Queue) String() string {
	el := this.front

	var stringBuilder string
	stringBuilder = fmt.Sprintln("BEGIN: Queue elements")
	for el != nil {
		stringBuilder = fmt.Sprintln(stringBuilder, el.value.value)
		el = el.next
	}
	stringBuilder = fmt.Sprintln(stringBuilder, "END")
	return stringBuilder 
} 

// Queue Element is defining the structure of each element in the queue
type QueueElement struct {
	value *HeapTreeNode
	next *QueueElement
}

// Node type of a Max-Heap Data Structure
type HeapTreeNode struct {
	value int
	parent *HeapTreeNode
	left *HeapTreeNode
	right *HeapTreeNode
}

func (this *HeapTreeNode) displayTree() string {
	var stringBuilder string
	stringBuilder = fmt.Sprintf("\nBEGIN: Tree elements -> traversal breadth-first\n")

	traversalQueue := new(Queue)
	node := this
	for node != nil {
		stringBuilder = fmt.Sprintln(stringBuilder, "node: ", node.value, "parent", node.parent, "left", node.left, "right", node.right)

		// breadth-first traversal
		if node.left != nil { traversalQueue.enqueue(&QueueElement{ value:node.left }) }
		if node.right != nil { traversalQueue.enqueue(&QueueElement{ value:node.right }) }

		// get the next in queue
		nodeEl := traversalQueue.dequeue()
		if nodeEl == nil { break }

		// try to insert node
		node = nodeEl.value
	}

	stringBuilder = fmt.Sprintln(stringBuilder, "END")
	return stringBuilder 
}

// Method to get the number of children
func (this *HeapTreeNode) numberOfChildren() int {
	result := 0
	if this.left != nil { result++ }
	if this.right != nil { result++ }
	return result
}

func (this *HeapTreeNode) addChild(val int) *HeapTreeNode {
	if this.numberOfChildren() == 2 {
		return nil
	}

	if this.left == nil {
		this.left = &HeapTreeNode{ value: val, parent: this }
		return this.left
	} else {
		this.right = &HeapTreeNode{ value: val, parent: this }
		return this.right
	}
}

func (this *HeapTreeNode) setChildLeft(node *HeapTreeNode){
	this.left = node
	if node != nil { this.left.parent = this }
}

func (this *HeapTreeNode) setChildRight(node *HeapTreeNode){
	this.right = node
	if node != nil { this.right.parent = this }
}

func (this *HeapTreeNode) swap(other *HeapTreeNode) {
	this.value, other.value = other.value, this.value
}

func (this *HeapTreeNode) removeRootAndRearrange2() *HeapTreeNode {

	var maxUInt = ^uint(0)
	var maxInt = int(maxUInt >> 1)
	var minInt = ^maxInt

	if this.numberOfChildren() == 0 {
		return nil
	}
	
	// set the value as minInt and then swap it until it arrives at the leaves level and then delete it
	this.value = minInt

	node := this
	for node.numberOfChildren() > 0 {
		
		if node.numberOfChildren() == 1 {
			if node.left != nil { node.swap(node.left); node = node.left }
			if node.right != nil { node.swap(node.right); node = node.right }
		} else if node.numberOfChildren() == 2{
			if node.left.value > node.right.value {
				node.swap(node.left) 
				node = node.left
			} else {
				node.swap(node.right) 
				node = node.right
			}
		} 
	}

	// remove the links of the node
	if node == node.parent.left { node.parent.left = nil }
	if node == node.parent.right { node.parent.right = nil }
	node.parent = nil
	
	return this
}

func (this *HeapTreeNode) removeRootAndRearrange() *HeapTreeNode {
	if this.numberOfChildren() == 0 {
		return nil
	}
	if (this.numberOfChildren() == 1) {
		if this.left != nil { this.left.parent = nil; return this.left }
		if this.right != nil { this.right.parent = nil; return this.right }
	}

	var rearrangeNode func(*HeapTreeNode) *HeapTreeNode
	rearrangeNode = func(node *HeapTreeNode) *HeapTreeNode {
		if node.numberOfChildren() == 0 { return nil }
		if node.numberOfChildren() == 1 {
			if node.left != nil { return node.left }
			if node.right != nil { return node.right }
		}
		
		var newRoot *HeapTreeNode
		if node.left.value > node.right.value {
			newRoot = node.left
			newRoot.setChildLeft(rearrangeNode(newRoot))
			newRoot.setChildRight(node.right)
		} else {
			newRoot = node.right
			newRoot.setChildRight(rearrangeNode(newRoot))
			newRoot.setChildLeft(node.left)
		}

		return newRoot
	}

	newTree := rearrangeNode(this)
	newTree.parent = nil
	return newTree
}

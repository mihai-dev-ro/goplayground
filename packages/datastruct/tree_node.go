package datastruct

import "fmt"

// Constants defining the position of the children in node of the tree
type PositionChild int

const (
	None PositionChild = iota
	Left
	Right
)

// TreeNode is defining a node in a tree data structure
type TreeNode struct {
	parent     *TreeNode
	childLeft  *TreeNode
	childRight *TreeNode
	Value      int
}

// implement method of stringer interface
func (currentNode *TreeNode) String() string {
	if currentNode == nil {
		return "nil"
	}

	return fmt.Sprintf("%v", currentNode.Value)
}

// DisplayTree traverse the tree in breadth-first manner and prints it
func (currentNode *TreeNode) DisplayTree() string {
	var stringBuilder string
	stringBuilder = fmt.Sprintf("\nBEGIN: Tree elements -> traversal breadth-first\n")

	traversalQueue := new(Queue)
	node := currentNode
	ok := true

	for &node != nil {
		stringBuilder = fmt.Sprintln(
			stringBuilder,
			"node: ", node,
			"| parent:", node.GetParent(),
			"| left", node.GetChildLeft(),
			"| right:", node.GetChildRight(),
		)

		// breadth-first traversal
		nodeLeft := node.GetChildLeft()
		if nodeLeft != nil {
			traversalQueue.Enqueue(nodeLeft)
		}
		nodeRight := node.GetChildRight()
		if nodeRight != nil {
			traversalQueue.Enqueue(nodeRight)
		}

		// get the next in queue
		nodeInQueue := traversalQueue.Dequeue()
		if nodeInQueue == nil {
			break
		}

		// try to extract node from queue
		node, ok = nodeInQueue.(*TreeNode)
		if !ok {
			panic("Value stored in Queue is not convertible to *TreeNode")
		}
	}

	stringBuilder = fmt.Sprintln(stringBuilder, "END")
	return stringBuilder
}

// NumberOfChildren returns the number of children
func (currentNode *TreeNode) NumberOfChildren() int {
	result := 0
	if currentNode.childLeft != nil {
		result++
	}

	if currentNode.childRight != nil {
		result++
	}

	return result
}

// AddChild creates a new node with the value send as argument and links it to the currentNode
// by assigning to the next available position (Left or Right)
func (currentNode *TreeNode) AddChild(val int, position PositionChild) *TreeNode {
	if currentNode.NumberOfChildren() == 2 {
		return nil
	}

	nodeInserted := &TreeNode{Value: val, parent: currentNode}
	switch position {
	case Left:
		currentNode.childLeft = nodeInserted
	case Right:
		currentNode.childRight = nodeInserted
	default: // next available
		if currentNode.childLeft == nil {
			currentNode.childLeft = nodeInserted
		} else {
			currentNode.childRight = nodeInserted
		}
	}

	return nodeInserted
}

// RemoveChild removes one of the child
func (currentNode *TreeNode) RemoveChild(node *TreeNode) {
	if node == nil {
		return
	}

	if currentNode.childLeft == node {
		currentNode.childLeft.parent = nil
		currentNode.childLeft = nil
	} else {
		currentNode.childRight.parent = nil
		currentNode.childRight = nil
	}
}

// ClearChildren removes all children
func (currentNode *TreeNode) ClearChildren() {
	currentNode.RemoveChild(currentNode.childLeft)
	currentNode.RemoveChild(currentNode.childRight)
}

// SetParent assigns the parent of the current node
func (currentNode *TreeNode) SetParent(node *TreeNode) {
	if currentNode.parent != nil {
		currentNode.parent.RemoveChild(currentNode)
	}

	currentNode.parent = node
}

// GetParent returns the parent
func (currentNode TreeNode) GetParent() *TreeNode {
	return currentNode.parent
}

// GetChildLeft returns the left child
func (currentNode TreeNode) GetChildLeft() *TreeNode {
	return currentNode.childLeft
}

// GetChildRight returns the right child
func (currentNode TreeNode) GetChildRight() *TreeNode {
	return currentNode.childRight
}

// SwapValue interchanges the values between 2 nodes
func (currentNode *TreeNode) SwapValue(other *TreeNode) {
	currentNode.Value, other.Value = other.Value, currentNode.Value
}

// TraverseBreadthFirst the tree starting with the current node and call the handler with each node
func (currentNode *TreeNode) TraverseBreadthFirst(
	handler func(*TreeNode) ([]*TreeNode, bool),
) {
	traversalQueue := new(Queue)

	node := currentNode
	ok := true

	for node != nil {

		if nodes, ok := handler(node); ok {
			for _, n := range nodes {
				traversalQueue.Enqueue(n)
			}
		} else {
			break
		}

		// get the next in queue
		nodeInQueue := traversalQueue.Dequeue()
		if nodeInQueue == nil {
			break
		}

		// try to extract node from queue
		node, ok = nodeInQueue.(*TreeNode)
		if !ok {
			panic("Value stored in Queue is not convertible to *TreeNode")
		}

	}
}

// TraverseDepthFirst the tree starting with the current node and call the handler with each node
func (currentNode *TreeNode) TraverseDepthFirst(
	visitor func(*TreeNode),
) {
	traversalStack := new(Stack)

	node := currentNode
	ok := true

	var trackNode func(*TreeNode)
	trackNode = func(node *TreeNode) {
		if node.GetChildRight() != nil {
			trackNode(node.GetChildRight())
		}
		traversalStack.Push(node)
		if node.GetChildLeft() != nil {
			trackNode(node.GetChildLeft())
		}
	}
	trackNode(currentNode) // create the stack

	for traversalStack.Len() > 0 {

		// get the next in stack
		nodeInStack := traversalStack.Pop()
		if nodeInStack == nil {
			return
		}

		// try to extract node from queue
		node, ok = nodeInStack.(*TreeNode)
		if !ok {
			panic("Value stored in Stack is not convertible to *TreeNode")
		}

		visitor(node)
	}
}

// InOrderVisit a binary search tree
func (currentNode *TreeNode) InOrderVisit(
	visitor func(*TreeNode),
) {
	var visitNode func(*TreeNode)
	visitNode = func(node *TreeNode) {
		if node.GetChildLeft() != nil {
			visitNode(node.GetChildLeft())
		}
		visitor(node)
		if node.GetChildRight() != nil {
			visitNode(node.GetChildRight())
		}
	}
	visitNode(currentNode)
}

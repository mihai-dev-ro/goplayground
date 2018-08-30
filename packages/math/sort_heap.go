package math

import (
	d "exercises/packages/datastruct"
)

// SortHeap implements the Heap Sort algorithm that implies building a heap data structure (which is a tree that
// has the nodes arranged in a certain/chosen order: max-heap or min-heap. A max-heap tree has the property
// that each parent node has a greater value that each of child nodes) and then progressively moving the
// root of the tree to append/prepend the sorted list.
// After each move of the root of the tree to the sorted list, we must re-arrange the heap data structure so
// that it remains a valid heap
//
// Complexity analysis
//    - time complexity: O(n*log(n))
//    - space complexity:O(n)
func SortHeap(list []int) []int {
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
		result[index] = tree.Value
		index--
		// tree = tree.removeRootAndRearrange2()
		tree.TraverseBreadthFirst(removeRootAndRearrange)
		// fmt.Println(tree.DisplayTree())
	}

	return result
}

func heapifyList(list []int) *d.TreeNode {

	var treeRoot *d.TreeNode

	var availableParentsQueue = new(d.Queue)
	var availableParent *d.TreeNode
	ok := true

	for _, val := range list {
		if treeRoot == nil {
			treeRoot = &d.TreeNode{Value: val}
			availableParentsQueue.Enqueue(treeRoot)
			continue
		} else {
			if availableParent == nil || availableParent.NumberOfChildren() == 2 {
				availableParent, ok = availableParentsQueue.Dequeue().(*d.TreeNode)
				if !ok {
					panic("ElementQueue should hold a TreeNode value")
				}
			}
			newNode := insertNode(availableParent, val)
			// new nodes added to the queue in the order they are added
			availableParentsQueue.Enqueue(newNode)
		}
	}

	return treeRoot
}

func insertNode(parentNode *d.TreeNode, val int) *d.TreeNode {

	if parentNode == nil {
		panic("Parent node should never be empty. It is the liaison required by the new node to be created")
	}

	// add the new value to the data structure
	var insertedNode = parentNode.AddChild(val, d.None)

	// 2. re-arrange nodes so that the data structure remains a valid max-heap data structure
	isSwapPerformed := true
	node := insertedNode
	for isSwapPerformed && node.GetParent() != nil {

		if node.Value > node.GetParent().Value {
			node.SwapValue(node.GetParent())
			isSwapPerformed = true
			node = node.GetParent()
		} else {
			isSwapPerformed = false
		}
	}

	return insertedNode
}

func compareNodes(i *d.TreeNode, j *d.TreeNode) int {
	if i.Value > j.Value {
		return 1
	} else if i.Value < j.Value {
		return -1
	}

	return 0
}

func removeRootAndRearrange(
	currentNode *d.TreeNode,
) (swappedNodes []*d.TreeNode, shouldContinue bool) {

	if currentNode.NumberOfChildren() == 0 {
		// remove it
		currentNode.SetParent(nil)

		// return end of action
		swappedNodes = nil
		shouldContinue = false
		return
	}

	swappedNodes = make([]*d.TreeNode, 1)
	if currentNode.NumberOfChildren() == 1 {
		if currentNode.GetChildLeft() != nil {
			swappedNodes[0] = currentNode.GetChildLeft()
		} else {
			swappedNodes[0] = currentNode.GetChildRight()
		}
	} else {
		if currentNode.GetChildLeft().Value >= currentNode.GetChildRight().Value {
			swappedNodes[0] = currentNode.GetChildLeft()
		} else {
			swappedNodes[0] = currentNode.GetChildRight()
		}
	}

	currentNode.SwapValue(swappedNodes[0])
	shouldContinue = true
	return
}

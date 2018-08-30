package math

import (
	d "exercises/packages/datastruct"
)

// SortBinary implements the binary sorting algorithm which requires the construct of a binary tree
// where each subtree is being organised is such a way that any child nodes with a lesser value
// are laid to the left and those with greater value are laid to the right
//
// Complexity analysis
//    - time complexity: O(n*log(n)) as average and pesimist is O(n*n)
//    - space complexity: O(n)
func SortBinary(list []int) []int {
	// empty or 1-item list; just return it as there is nothing to sort
	if len(list) < 2 {
		return list
	}

	// build the binary search data tree
	var searchTree *d.TreeNode
	for _, val := range list {
		if searchTree == nil {
			searchTree = &d.TreeNode{Value: val}
			continue
		}
		insertValue(searchTree, val)
	}

	// traverse breadth first and populate the sorted list
	result := make([]int, len(list))
	i := 0
	extractValue := func(node *d.TreeNode) {
		result[i] = node.Value
		i++
	}
	searchTree.TraverseDepthFirst(extractValue)

	return result
}

func insertValue(tree *d.TreeNode, value int) {
	if tree == nil {
		return
	}

	node := tree
	shouldContinue := true

	for shouldContinue {
		if node.Value >= value {
			nodeLeft := node.GetChildLeft()
			if nodeLeft == nil {
				// insert new node here
				node.AddChild(value, d.Left)
				shouldContinue = false
			} else {
				node = nodeLeft
			}
		} else {
			nodeRight := node.GetChildRight()
			if nodeRight == nil {
				// insert new node here
				node.AddChild(value, d.Right)
				shouldContinue = false
			} else {
				node = nodeRight
			}
		}
	}
}

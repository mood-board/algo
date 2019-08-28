// Problem:
// Given a binary tree, determine if it is "superbalanced" - the difference
// between the depths of any two leaf nodes is no greater than 1.
//
// Example:
// Input:
//          1
//       2     3
//     4   5      7
//           6  8   9
//                    10
// Output: false
// Even though this tree is balanced by definition, it is not "superbalanced".
//
// Solution:
// Use a depth-first walk through the tree and keep track of the depth as we
// go.
// Every time we found a leaf with a new depth, there are two ways that the
// tree could be unbalanced:
// - There are more than 2 different leaf depths.
// - There are exactly 2 depths but they are more than 1 apart.
//
// Cost:
// O(n) time, O(n) space.
// The worst case is that we have to iterate all nodes in the tree so the time
// complexity is O(n). For space complexity, we have to keep track of the all
// the nodes at every depth. Hence, it is O(n).

package main

import (
	"math"
	"reflect"
	"testing"
)

func TestIsSuperBalanced(t *testing.T) {
	// define test cases' input.
	t1 := &BinaryTree{nil, 1, nil}

	t2 := &BinaryTree{nil, 1, nil}
	t2.right = &BinaryTree{nil, 2, nil}

	t3 := &BinaryTree{nil, 1, nil}
	t3.right = &BinaryTree{nil, 2, nil}
	t3.right.right = &BinaryTree{nil, 3, nil}

	t4 := &BinaryTree{nil, 1, nil}
	t4.left = &BinaryTree{nil, 2, nil}
	t4.right = &BinaryTree{nil, 3, nil}
	t4.right.right = &BinaryTree{nil, 4, nil}

	t5 := &BinaryTree{nil, 1, nil}
	t5.left = &BinaryTree{nil, 2, nil}
	t5.right = &BinaryTree{nil, 3, nil}
	t5.right.right = &BinaryTree{nil, 4, nil}
	t5.right.right.right = &BinaryTree{nil, 5, nil}

	t6 := &BinaryTree{nil, 1, nil}
	t6.left = &BinaryTree{nil, 2, nil}
	t6.left.left = &BinaryTree{nil, 4, nil}
	t6.left.right = &BinaryTree{nil, 5, nil}
	t6.left.right.right = &BinaryTree{nil, 6, nil}
	t6.right = &BinaryTree{nil, 3, nil}
	t6.right.right = &BinaryTree{nil, 7, nil}
	t6.right.right.left = &BinaryTree{nil, 8, nil}
	t6.right.right.right = &BinaryTree{nil, 9, nil}
	t6.right.right.right.right = &BinaryTree{nil, 10, nil}

	// define their outputs.
	tests := []struct {
		in       *BinaryTree
		expected bool
	}{
		{t1, true},
		{t2, true},
		{t3, true},
		{t4, true},
		{t5, false},
		{t6, false},
	}

	for _, tt := range tests {
		result := isSuperBalanced(tt.in)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("should be %v instead %v", tt.expected, result)
		}
	}
}

// BinaryTree represents a binary tree.
type BinaryTree struct {
	left  *BinaryTree
	value int
	right *BinaryTree
}

// treeDepth holds the tree and its depth level.
type treeDepth struct {
	tree  *BinaryTree
	depth int
}

func contains(s []int, target int) bool {
	for _, v := range s {
		if v == target {
			return true
		}
	}

	return false
}

func is1Apart(a, b int) bool {
	if math.Abs(float64(a)-float64(b)) > 1 {
		return true
	}

	return false
}

func isSuperBalanced(t *BinaryTree) bool {
	// return true if the tree has no leaf.
	if t == nil {
		return true
	}

	// depths holds a list of depth that we have seen.
	depths := []int{}

	// stack keeps track of the tree level and its depth.
	stack := []treeDepth{}
	stack = append(stack, treeDepth{t, 0})

	for len(stack) > 0 {
		// pop a tree and its depth from the top of our stack.
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// if we found a leaf, add the new depth to the list if we haven't seen it.
		if current.tree.left == nil && current.tree.right == nil {
			if !contains(depths, current.depth) {
				depths = append(depths, current.depth)
			}

			// short-circuit to determine if the tree is unbalanced:
			// - more than 2 different leaf depths
			// - 2 leaf depths that are more than 1 apart
			if (len(depths) > 2) || (len(depths) == 2 && is1Apart(depths[1], depths[0])) {
				return false
			}
		}

		// keep walking down the tree and keep track of the depth.
		if current.tree.left != nil {
			stack = append(stack, treeDepth{current.tree.left, current.depth + 1})
		}
		if current.tree.right != nil {
			stack = append(stack, treeDepth{current.tree.right, current.depth + 1})
		}
	}

	return true
}
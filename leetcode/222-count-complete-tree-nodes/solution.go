// Given the root of a complete binary tree, return the number of the nodes in the tree.
//
// According to Wikipedia, every level, except possibly the last, is completely filled in a complete binary tree, and all nodes in the last level are as far left as possible. It can have between 1 and 2h nodes inclusive at the last level h.
//
// Design an algorithm that runs in less than O(n) time complexity.
// Input: root = [1,2,3,4,5,6]
// Output: 6
// Input: root = []
// Output: 0
// Input: root = [1]
// Output: 1
//
// Constraints:
// The number of nodes in the tree is in the range [0, 5 * 104].
// 0 <= Node.val <= 5 * 104
// The tree is guaranteed to be complete.

package countcompletetreenodes

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// O(n) time, O(log n) space
// func countNodes(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	return countNodes(root.Left) + countNodes(root.Right) + 1
// }

// O(log^2 n) time, O(log n) space
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lh := leftHeight(root)
	rh := rightHeight(root)

	// perfect binary tree of height lh: 2^lh - 1 nodes
	if lh == rh {
		return int(math.Pow(2, float64(lh))) - 1
	}

	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

func leftHeight(node *TreeNode) int {
	h := 0
	for node != nil {
		h++
		node = node.Left
	}
	return h
}

func rightHeight(node *TreeNode) int {
	h := 0
	for node != nil {
		h++
		node = node.Right
	}
	return h
}

// TODO: revisit this problem

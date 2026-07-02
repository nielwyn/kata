package countcompletetreenodes

import "testing"

func TestCountNodes(t *testing.T) {
	tests := []struct {
		vals []int
		want int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 6},
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 7},
		{[]int{1, 2, 3, 4}, 4},
	}

	for _, tt := range tests {
		root := buildBinaryTree(tt.vals)
		got := countNodes(root)
		if got != tt.want {
			t.Errorf("countNodes(%v) = %d, want %d", tt.vals, got, tt.want)
		}
	}
}

// buildTree builds a complete binary tree from a level-order slice of values.
func buildBinaryTree(vals []int) *TreeNode {
	if len(vals) == 0 {
		return nil
	}
	nodes := make([]*TreeNode, len(vals))
	for i, v := range vals {
		nodes[i] = &TreeNode{Val: v}
	}
	for i := range nodes {
		left := 2*i + 1
		right := 2*i + 2
		if left < len(nodes) {
			nodes[i].Left = nodes[left]
		}
		if right < len(nodes) {
			nodes[i].Right = nodes[right]
		}
	}
	return nodes[0]
}

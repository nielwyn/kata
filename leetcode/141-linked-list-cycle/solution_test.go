package linkedlistcycle

import "testing"

func TestHasCycle(t *testing.T) {
	tests := []struct {
		vals []int
		pos  int
		want bool
	}{
		{[]int{3, 2, 0, -4}, 1, true},
		{[]int{1, 2}, -1, false},
		{[]int{1}, -1, false},
		{[]int{}, -1, false},
		{[]int{1}, 0, true},
		{[]int{1, 2}, 0, true},
		{[]int{1, 2, 3, 4, 5}, -1, false},
		{[]int{1, 2, 3, 4, 5}, 0, true},
		{[]int{1, 2, 3, 4, 5}, 4, true},
	}

	for _, tt := range tests {
		head := buildCyclicList(tt.vals, tt.pos)
		got := hasCycle(head)
		if got != tt.want {
			t.Errorf("hasCycle(%v, pos=%d) = %v, want %v", tt.vals, tt.pos, got, tt.want)
		}
	}
}

func buildCyclicList(vals []int, pos int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	nodes := make([]*ListNode, len(vals))
	for i, v := range vals {
		nodes[i] = &ListNode{Val: v}
	}
	for i, node := range nodes[:len(nodes)-1] {
		node.Next = nodes[i+1]
	}
	if pos >= 0 {
		nodes[len(nodes)-1].Next = nodes[pos]
	}
	return nodes[0]
}

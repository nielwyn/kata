package intersectionoftwolinkedlists

import (
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {
	tests := []struct {
		name   string
		vals1  []int
		vals2  []int
		shared []int
	}{
		{
			name:   "intersection in the middle",
			vals1:  []int{4, 1},
			vals2:  []int{5, 6, 1},
			shared: []int{8, 4, 5},
		},
		{
			name:   "no intersection",
			vals1:  []int{2, 6, 4},
			vals2:  []int{1, 5},
			shared: []int{},
		},
		{
			name:   "intersection at head of A",
			vals1:  []int{},
			vals2:  []int{3},
			shared: []int{7, 2, 4},
		},
		{
			name:   "both lists are entirely the shared part",
			vals1:  []int{},
			vals2:  []int{},
			shared: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shared := buildLinkedList(tt.shared)

			headA := buildLinkedList(tt.vals1)
			headB := buildLinkedList(tt.vals2)

			if tail := getTail(headA); tail != nil {
				tail.Next = shared
			} else {
				headA = shared
			}

			if tail := getTail(headB); tail != nil {
				tail.Next = shared
			} else {
				headB = shared
			}

			got := getIntersectionNode(headA, headB)
			if got != shared {
				t.Errorf("got %v, want %v", got, shared)
			}
		})
	}
}

func buildLinkedList(vals []int) *ListNode {
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
	return nodes[0]
}

func getTail(n *ListNode) *ListNode {
	if n == nil {
		return nil
	}
	for n.Next != nil {
		n = n.Next
	}
	return n
}

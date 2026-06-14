package intersectionoftwolinkedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pA := headA
	pB := headB
	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}

		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}
	return pA
}

// Hash set
// func getIntersectionNodeHashSet(headA, headB *ListNode) *ListNode {
// 	visited := make(map[*ListNode]bool)
// 	for node := headA; node != nil; node = node.Next {
// 		visited[node] = true
// 	}
// 	for node := headB; node != nil; node = node.Next {
// 		if visited[node] {
// 			return node
// 		}
// 	}
// 	return nil
// }

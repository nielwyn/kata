// Given head, the head of a linked list, determine if the linked list has a cycle in it.
//
// There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.
//
// Return true if there is a cycle in the linked list. Otherwise, return false.

// Input: head = [3,2,0,-4], pos = 1
// Output: true
// Input: head = [1], pos = -1
// Output: false

package linkedlistcycle

type ListNode struct {
	Val  int
	Next *ListNode
}

// Hash table
func hasCycle(head *ListNode) bool {
	checked := make(map[*ListNode]bool)
	node := head
	for node != nil {
		if checked[node] {
			return true
		}
		checked[node] = true
		node = node.Next
	}
	return false
}

// Floyd's Cycle Detection
// func hasCycle(head *ListNode) bool {
// 	slow, fast := head, head
//
// 	for fast != nil && fast.Next != nil {
// 		slow = slow.Next
// 		fast = fast.Next.Next
//
// 		if slow == fast {
// 			return true
// 		}
// 	}
//
// 	return false
// }

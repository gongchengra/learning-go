package main

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}
	fast := head
	for i := 0; i < k; i++ {
		if fast.Next == nil {
			return rotateRight(head, k%(i+1))
		}
		fast = fast.Next
	}
	slow := head
	for fast.Next != nil {
		slow, fast = slow.Next, fast.Next
	}
	newHead := slow.Next
	slow.Next, fast.Next = nil, head
	return newHead
}

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func s2l(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	res := &ListNode{Val: nums[0]}
	tmp := res
	for i := 1; i < len(nums); i++ {
		tmp.Next = &ListNode{Val: nums[i]}
		tmp = tmp.Next
	}
	return res
}

func l2s(head *ListNode) (res []int) {
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	lessHead, moreHead := &ListNode{}, &ListNode{}
	lessEnd, moreEnd := lessHead, moreHead
	for ; head != nil; head = head.Next {
		if head.Val < x {
			lessEnd.Next = head
			lessEnd = lessEnd.Next
		} else {
			moreEnd.Next = head
			moreEnd = moreEnd.Next
		}
	}
	lessEnd.Next, moreEnd.Next = moreHead.Next, nil
	return lessHead.Next
}

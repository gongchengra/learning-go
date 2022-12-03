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

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	tmp := head
	for tmp.Next != nil {
		if tmp.Val == tmp.Next.Val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
	return head
}

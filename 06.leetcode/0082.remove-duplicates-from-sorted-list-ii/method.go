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
	if head == nil || head.Next == nil {
		return head
	}
	if head.Val == head.Next.Val {
		val := head.Val
		head = head.Next.Next
		for head != nil && head.Val == val {
			head = head.Next
		}
		return deleteDuplicates(head)
	}
	head.Next = deleteDuplicates(head.Next)
	return head
}

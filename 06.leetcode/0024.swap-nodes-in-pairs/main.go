package main

import "fmt"

func main() {
	fmt.Println(l2s(swapPairs(s2l([]int{1, 2, 3, 4}))))
	// fmt.Println(l2s(swapPairs(s2l([]int{}))))
	// fmt.Println(l2s(swapPairs(s2l([]int{1}))))
}

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

// method 1:
/*
func swapPairs(head *ListNode) *ListNode {
	tmp := &ListNode{Next: head}
	for pt := tmp; pt != nil && pt.Next != nil && pt.Next.Next != nil; {
		pt, pt.Next, pt.Next.Next, pt.Next.Next.Next = pt.Next, pt.Next.Next, pt.Next.Next.Next, pt.Next
	}
	return tmp.Next
}
*/

// method 2
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newhead := head.Next
	head.Next = swapPairs(newhead.Next)
	newhead.Next = head
	return newhead
}

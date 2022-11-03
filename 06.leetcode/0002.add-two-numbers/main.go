package main

import "fmt"

func main() {
	fmt.Println(list2int(
		addTwoNumbers(
			int2list([]int{0}),
			int2list([]int{0}))))
	fmt.Println(list2int(
		addTwoNumbers(
			int2list([]int{2, 4, 3}),
			int2list([]int{5, 6, 4}))))
	fmt.Println(list2int(
		addTwoNumbers(
			int2list([]int{9, 9, 9, 9, 9, 9, 9}),
			int2list([]int{9, 9, 9, 9}))))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func list2int(head *ListNode) (res []int) {
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return
}

func int2list(nums []int) (res *ListNode) {
	if len(nums) == 0 {
		return nil
	}
	res = &ListNode{}
	tmp := res
	for _, v := range nums {
		tmp.Next = &ListNode{Val: v}
		tmp = tmp.Next
	}
	return res.Next
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		l1 = &ListNode{}
	}
	if l2 == nil {
		l2 = &ListNode{}
	}
	res := &ListNode{}
	res.Val = l1.Val + l2.Val
	if res.Val > 9 {
		res.Val -= 10
		if l2.Next == nil {
			l2.Next = &ListNode{}
		}
		l2.Next.Val += 1
	}
	if l1.Next != nil || l2.Next != nil {
		res.Next = addTwoNumbers(l1.Next, l2.Next)
	}
	return res
}

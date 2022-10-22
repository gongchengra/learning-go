package main

import "fmt"

func main() {
	fmt.Println(l2s(mergeKLists(ss2ls(
		[][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}))))
	fmt.Println(l2s(mergeKLists(ss2ls(
		[][]int{{1, 4, 5}, {1, 3, 4}, {3, 6, 9}}))))
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

func ss2ls(numss [][]int) []*ListNode {
	res := []*ListNode{}
	for _, nums := range numss {
		res = append(res, s2l(nums))
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

func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}
	num := length / 2
	left := mergeKLists(lists[:num])
	right := mergeKLists(lists[num:])
	return mergeTwoLists(left, right)
}

func mergeTwoLists(l, r *ListNode) *ListNode {
	if l == nil {
		return r
	}
	if r == nil {
		return l
	}
	if l.Val < r.Val {
		l.Next = mergeTwoLists(l.Next, r)
		return l
	}
	r.Next = mergeTwoLists(l, r.Next)
	return r
}

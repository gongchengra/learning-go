package main

import "fmt"

func main() {
	//     fmt.Println(l2s(rotateRight(s2l([]int{1, 2, 3, 4, 5}), 2)))
	fmt.Println(l2s(rotateRight(s2l([]int{0, 1, 2}), 4)))
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

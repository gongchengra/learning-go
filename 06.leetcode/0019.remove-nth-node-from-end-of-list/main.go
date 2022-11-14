package main

import "fmt"

func main() {
	fmt.Println(l2s(removeNthFromEnd(s2l([]int{1, 2, 3, 4, 5}), 2)))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	d, headIsNthFromEnd := getDaddy(head, n)
	if headIsNthFromEnd {
		return head.Next
	}
	d.Next = d.Next.Next
	return head
}

func getDaddy(head *ListNode, n int) (daddy *ListNode, headIsNthFromEnd bool) {
	daddy = head
	for head != nil {
		if n < 0 {
			daddy = daddy.Next
		}
		n--
		head = head.Next
	}
	headIsNthFromEnd = (n == 0)
	return
}

// convert *ListNode to []int
func l2s(head *ListNode) []int {
	res := []int{}

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

// convert []int to *ListNode
func s2l(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	res := &ListNode{
		Val: nums[0],
	}
	temp := res
	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{
			Val: nums[i],
		}
		temp = temp.Next
	}

	return res
}

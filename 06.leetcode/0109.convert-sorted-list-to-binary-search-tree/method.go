package main

func sortedListToBST(head *ListNode) *TreeNode {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return sortedArrayToBST(res)
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	return &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}
}

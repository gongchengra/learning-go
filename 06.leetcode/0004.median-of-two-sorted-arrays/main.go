package main

import "fmt"

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2, l := len(nums1), len(nums2), len(nums1)+len(nums2)
	i, j := 0, 0
	res := make([]int, l1+l2)
	for n := 0; n < l1+l2; n++ {
		if i == l1 || (i < l1 && j < l2 && nums1[i] > nums2[j]) {
			res[n] = nums2[j]
			j++
			continue
		}
		if j == l2 || (i < l1 && j < l2 && nums1[i] <= nums2[j]) {
			res[n] = nums1[i]
			i++
		}
	}
	if l%2 == 0 {
		return float64(res[l/2]+res[l/2-1]) / 2.0
	} else {
		return float64(res[l/2])
	}
}

package main

import (
	"fmt"
	"laogongshuo.com/search/BinarySearch"
	"laogongshuo.com/search/InterpolationSearch"
	"laogongshuo.com/search/JumpSearch"
	"laogongshuo.com/search/LinearSearch"
	"math/rand"
	"sort"
	"time"
)

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([]int, random.Intn(100-10)+10)
	for i := range arr {
		arr[i] = random.Intn(100)
	}
	find := arr[len(arr)/2]
	sort.Ints(arr)
	fmt.Println(arr)
	//	find := random.Intn(100)
	fmt.Println(find)
	fmt.Println(BinarySearch.BinarySearch(arr, find))
	fmt.Println(InterpolationSearch.InterpolationSearch(arr, find))
	fmt.Println(JumpSearch.JumpSearch(arr, find))
	fmt.Println(LinearSearch.LinearSearch(arr, find))
}

package main

import (
	"fmt"
	"laogongshuo.com/search/BinarySearch"
	"laogongshuo.com/search/InterpolationSearch"
	"laogongshuo.com/search/JumpSearch"
	"laogongshuo.com/search/LinearSearch"
	"laogongshuo.com/sort/BubbleSort"
	"laogongshuo.com/sort/CountingSort"
	"laogongshuo.com/sort/HeapSort"
	"laogongshuo.com/sort/InsertionSort"
	"laogongshuo.com/sort/MergeSort"
	"laogongshuo.com/sort/QuickSort"
	"laogongshuo.com/sort/RadixSort"
	"laogongshuo.com/sort/SelectionSort"
	"laogongshuo.com/sort/ShellSort"
	"math/rand"
	"time"
)

type Heap struct{}

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([]int, random.Intn(20)+10)
	for i := range arr {
		arr[i] = random.Intn(100)
	}
	fmt.Println("arr:")
	fmt.Println(arr)
	find := arr[len(arr)/2]
	fmt.Println("find:")
	fmt.Println(find)
	arrBubble := make([]int, len(arr))
	copy(arrBubble, arr)
	BubbleSort.BubbleSort(arrBubble)
	fmt.Println("arrBubble:")
	fmt.Println(arrBubble)
	fmt.Println(BinarySearch.BinarySearch(arrBubble, find))
	arrCounting := append([]int{}, arr...)
	CountingSort.CountingSort(arrCounting)
	fmt.Println("arrCounting:")
	fmt.Println(arrCounting)
	fmt.Println(InterpolationSearch.InterpolationSearch(arrCounting, find))
	arrHeap := append([]int{}, arr...)
	var heap = new(HeapSort.Heap)
	heap.HeapSort(arrHeap)
	fmt.Println("arrHeap:")
	fmt.Println(arrHeap)
	fmt.Println(JumpSearch.JumpSearch(arrHeap, find))
	arrInsertion := append([]int{}, arr...)
	InsertionSort.InsertionSort(arrInsertion)
	fmt.Println("arrInsertion:")
	fmt.Println(arrInsertion)
	fmt.Println(LinearSearch.LinearSearch(arrInsertion, find))
	arrMerge := append([]int{}, arr...)
	arrMerge = MergeSort.MergeSort(arr)
	fmt.Println("arrMerge:")
	fmt.Println(arrMerge)
	arrQuick := append([]int{}, arr...)
	QuickSort.QuickSort(arrQuick, 0, (len(arr) - 1))
	fmt.Println("arrQuick:")
	fmt.Println(arrQuick)
	arrRadix := append([]int{}, arr...)
	RadixSort.RadixSort(arrRadix, len(arr))
	fmt.Println("arrRadix:")
	fmt.Println(arrRadix)
	arrSelection := append([]int{}, arr...)
	SelectionSort.SelectionSort(arrSelection)
	fmt.Println("arrSelection:")
	fmt.Println(arrSelection)
	arrShell := append([]int{}, arr...)
	arrShell = ShellSort.ShellSort(arrShell)
	fmt.Println("arrShell:")
	fmt.Println(arrShell)
}

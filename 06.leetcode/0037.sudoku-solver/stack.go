package main

import (
	"container/list"
	"fmt"
)

func main() {
	stack := list.New()
	fmt.Println(stack.Len())
	b3 := [][]byte{
		[]byte("4.....8.5"),
		[]byte(".3......."),
		[]byte("...7....."),
		[]byte(".2.....6."),
		[]byte("....8.4.."),
		[]byte("....1...."),
		[]byte("...6.3.7."),
		[]byte("5..2....."),
		[]byte("1.4......"),
	}
	pos := calculatePossibility(b3)
	update(pos)
	printPos(pos)
	s := status(pos)
	fmt.Println(s)

	k, v := leastUnknow(pos)
	fmt.Println("unsolved", k, string(v))
	stack.PushBack(pos)
	pos[k] = []byte{v[0]}
	kv := make([]interface{}, 2)
	kv[0], kv[1] = k, remove(v, v[0])
	fmt.Println(kv[0].(int), string(kv[1].([]byte)))
	update(pos)
	printPos(pos)

	posSlice := [][][]byte{}
	for i := 0; i < 9; i++ {
		k, v = leastUnknow(pos)
		fmt.Println("unsolved", k, string(v))
		pos[k] = []byte{v[0]}
		kv[0], kv[1] = k, pos[k]
		fmt.Println(kv[0].(int), string(kv[1].([]byte)))
		update(pos)
		tmp := deepcopy(pos)
		stack.PushBack(tmp)
		posSlice = append(posSlice, tmp)
		printPos(pos)
	}

	fmt.Println("Pop")
	for i := len(posSlice) - 1; i >= 0; i-- {
		//         printPos(posSlice[i])
		tmp := stack.Remove(stack.Back()).([][]byte)
		printPos(tmp)
	}

}

func deepcopy(pos [][]byte) (res [][]byte) {
	for i, k := range pos {
		res = append(res, []byte{})
		for _, v := range k {
			res[i] = append(res[i], v)
		}
	}
	return
}

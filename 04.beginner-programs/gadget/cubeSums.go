package main

import (
	"fmt"
	"sort"
)

func main() {
	// 创建一个map来保存(a^3 + b^3)的结果和对应的a，b对
	cubeSums := make(map[int][][2]int)

	// 计算所有的a^3 + b^3，并把它们映射到相应的a，b对
	for a := 1; a < 1000; a++ {
		for b := a; b < 1000; b++ {
			sum := a*a*a + b*b*b
			cubeSums[sum] = append(cubeSums[sum], [2]int{a, b})
		}
	}

	type quad struct {
		a, b, c, d int
	}

	// 收集所有的组合
	var quads []quad

	// 遍历map找出所有c^3 + d^3有相同结果的解
	for _, pairs := range cubeSums {
		if len(pairs) > 1 {
			for i := 0; i < len(pairs); i++ {
				for j := i + 1; j < len(pairs); j++ {
					// 添加不重复的组合
					quads = append(quads, quad{pairs[i][0], pairs[i][1], pairs[j][0], pairs[j][1]})
				}
			}
		}
	}

	// 对所有组合进行全局排序
	sort.Slice(quads, func(i, j int) bool {
		if quads[i].a == quads[j].a {
			if quads[i].b == quads[j].b {
				if quads[i].c == quads[j].c {
					return quads[i].d < quads[j].d
				}
				return quads[i].c < quads[j].c
			}
			return quads[i].b < quads[j].b
		}
		return quads[i].a < quads[j].a
	})

	// 打印所有排序后的组合
	for _, q := range quads {
		fmt.Printf("a=%d, b=%d, c=%d, d=%d\n", q.a, q.b, q.c, q.d)
	}
}

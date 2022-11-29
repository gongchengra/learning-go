package main

func setZeroes(m [][]int) {
	r := make([]bool, len(m))
	c := make([]bool, len(m[0]))
	for rk, rv := range m {
		for ck, cv := range rv {
			if cv == 0 {
				r[rk] = true
				c[ck] = true
			}
		}
	}
	for i := range r {
		if r[i] {
			for j := range m[i] {
				m[i][j] = 0
			}
		}
	}
	for i := range c {
		if c[i] {
			for j := range m {
				m[j][i] = 0
			}
		}
	}
}

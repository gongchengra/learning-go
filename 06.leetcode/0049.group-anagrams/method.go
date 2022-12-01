package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	record := make(map[string][]string)
	for _, str := range strs {
		tmp := sortString(str)
		record[tmp] = append(record[tmp], str)
	}
	for _, v := range record {
		//         sort.Strings(v)
		res = append(res, v)
	}
	return res
}

func sortString(s string) string {
	bytes := []byte(s)
	tmp := make([]int, len(bytes))
	for i, b := range bytes {
		tmp[i] = int(b)
	}
	sort.Ints(tmp)
	for i, v := range tmp {
		bytes[i] = byte(v)
	}
	return string(bytes)
}

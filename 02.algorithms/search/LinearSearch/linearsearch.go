package LinearSearch

func LinearSearch(arr []int, query int) int {
	for i, val := range arr {
		if val == query {
			return i
		}
	}
	return -1
}

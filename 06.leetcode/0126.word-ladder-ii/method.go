package main

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	words := deleteBeginWord(wordList, beginWord)
	trans := map[string][]string{}
	isTransedEndWord := false
	cnt := 1
	var bfs func([]string, []string)
	bfs = func(words, nodes []string) {
		cnt++
		newWords := make([]string, 0, len(words))
		newNodes := make([]string, 0, len(words))
		for _, w := range words {
			isTransed := false
			for _, n := range nodes {
				if isTransable(n, w) {
					trans[n] = append(trans[n], w)
					isTransed = true
				}
			}
			if isTransed {
				newNodes = append(newNodes, w)
				if w == endWord {
					isTransedEndWord = true
				}
			} else {
				newWords = append(newWords, w)
			}
		}
		if isTransedEndWord || len(newWords) == 0 || len(newNodes) == 0 {
			return
		}
		bfs(newWords, newNodes)
	}
	nodes := []string{beginWord}
	bfs(words, nodes)
	res := [][]string{}
	if !isTransedEndWord {
		return res
	}
	path := make([]string, cnt)
	path[0] = beginWord
	var dfs func(int)
	dfs = func(idx int) {
		if idx == cnt {
			if path[idx-1] == endWord {
				res = append(res, deepCopy(path))
			}
			return
		}
		prev := path[idx-1]
		for _, w := range trans[prev] {
			path[idx] = w
			dfs(idx + 1)
		}
	}
	dfs(1)
	return res
}

func deleteBeginWord(words []string, beginWord string) []string {
	i, size := 0, len(words)
	for ; i < size; i++ {
		if words[i] == beginWord {
			break
		}
	}
	if i == size {
		return words
	}
	words[i] = words[size-1]
	return words[:size-1]
}

func deepCopy(src []string) []string {
	tmp := make([]string, len(src))
	copy(tmp, src)
	return tmp
}

func isTransable(a, b string) bool {
	onceAgain := false
	for i := range a {
		if a[i] != b[i] {
			if onceAgain {
				return false
			}
			onceAgain = true
		}
	}
	return true
}

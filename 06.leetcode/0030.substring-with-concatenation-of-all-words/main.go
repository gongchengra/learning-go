package main

import (
	"fmt"
)

func main() {
	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}))
	fmt.Println(findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}))
	fmt.Println(findSubstring("pjzkrkevzztxductzzxmxsvwjkxpvukmfjywwetvfnujhweiybwvvsrfequzkhossmootkmyxgjgfordrpapjuunmqnxxdrqrfgkrsjqbszgiqlcfnrpjlcwdrvbumtotzylshdvccdmsqoadfrpsvnwpizlwszrtyclhgilklydbmfhuywotjmktnwrfvizvnmfvvqfiokkdprznnnjycttprkxpuykhmpchiksyucbmtabiqkisgbhxngmhezrrqvayfsxauampdpxtafniiwfvdufhtwajrbkxtjzqjnfocdhekumttuqwovfjrgulhekcpjszyynadxhnttgmnxkduqmmyhzfnjhducesctufqbumxbamalqudeibljgbspeotkgvddcwgxidaiqcvgwykhbysjzlzfbupkqunuqtraxrlptivshhbihtsigtpipguhbhctcvubnhqipncyxfjebdnjyetnlnvmuxhzsdahkrscewabejifmxombiamxvauuitoltyymsarqcuuoezcbqpdaprxmsrickwpgwpsoplhugbikbkotzrtqkscekkgwjycfnvwfgdzogjzjvpcvixnsqsxacfwndzvrwrycwxrcismdhqapoojegggkocyrdtkzmiekhxoppctytvphjynrhtcvxcobxbvtfjiwmduhzjokkbctweqtigwfhzorjlkpuuliaipbtfldinyetoybvugevwvhhhweejogrghllsouipabfafcxnhukcbtmxzshoyyufjhzadhrelweszbfgwpkzlwxkogyogutscvuhcllphshivnoteztpxsaoaacgxyaztuixhunrowzljqfqrahosheukhahhbiaxqzfmmwcjxountkevsvpbzjnilwpoermxrtlfroqoclexxisrdhvfsindffslyekrzwzqkpeocilatftymodgztjgybtyheqgcpwogdcjlnlesefgvimwbxcbzvaibspdjnrpqtyeilkcspknyylbwndvkffmzuriilxagyerjptbgeqgebiaqnvdubrtxibhvakcyotkfonmseszhczapxdlauexehhaireihxsplgdgmxfvaevrbadbwjbdrkfbbc", []string{"dhvf", "sind", "ffsl", "yekr", "zwzq", "kpeo", "cila", "tfty", "modg", "ztjg", "ybty", "heqg", "cpwo", "gdcj", "lnle", "sefg", "vimw", "bxcb"}))
}

func findSubstring(s string, words []string) (res []int) {
	if len(words) == 0 || len(s) == 0 {
		return res
	}
	l, size := len(words[0]), len(words[0])*len(words)
	if size > len(s) {
		return res
	}
	wordlist := map[string]int{}
	for _, v := range words {
		wordlist[v]++
	}
	for i := 0; i <= len(s)-size; i++ {
		tmp := s[i : i+size]
		tmplist := map[string]int{}
		for k, v := range wordlist {
			tmplist[k] = v
		}
		for j := 0; j < size; j += l {
			idx, ok := tmplist[tmp[j:j+l]]
			if ok == false {
				break
			}
			if idx > 0 {
				tmplist[tmp[j:j+l]]--
			}
		}
		s := 0
		for _, v := range tmplist {
			s += v
		}
		if s == 0 {
			res = append(res, i)
		}
	}
	return res
}

/* method 2 use permutation as array

func findSubstring(s string, words []string) (res []int) {
	wordlist := map[string]int{}
	wordlen := len(strings.Join(words, ""))
	for _, arr := range permutations(words) {
		wordlist[strings.Join(arr, "")] = 1
	}
	for i := 0; i <= len(s)-wordlen; i++ {
		tmp := s[i : i+wordlen]
		idx, _ := wordlist[tmp]
		if idx == 1 {
			res = append(res, i)
		}
	}
	return res
}

func permutations(arr []string) [][]string {
	var helper func([]string, int)
	res := [][]string{}
	helper = func(arr []string, n int) {
		if n == 1 {
			tmp := make([]string, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
*/

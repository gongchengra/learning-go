package main

func fullJustify(words []string, maxWidth int) []string {
	space := ""
	for i := 0; i < maxWidth; i++ {
		space += " "
	}
	if len(words) == 0 {
		return []string{space}
	}
	ret, row := []string{}, []string{}
	row = append(row, words[0])
	line, size := len(words[0]), len(words[0])
	for i := 1; i < len(words); i++ {
		size = len(words[i])
		if line+size+1 > maxWidth {
			line -= len(row) - 1
			numspace, perspace, more := (maxWidth - line), 0, 0
			if len(row) > 1 {
				perspace = int(numspace / (len(row) - 1))
				more = numspace - perspace*(len(row)-1)
			} else {
				perspace = numspace
			}
			str := ""
			for j := 0; j < len(row); j++ {
				str += row[j]
				if j == 0 || j != len(row)-1 {
					if j < more {
						str += space[:perspace+1]
					} else {
						str += space[:perspace]
					}
				}
			}
			ret = append(ret, str)
			row = []string{words[i]}
			line = len(words[i])
		} else {
			row = append(row, words[i])
			line += size + 1
		}
	}
	str := ""
	for j := 0; j < len(row); j++ {
		str += row[j]
		if j != len(row)-1 {
			str += " "
		}
	}
	str += space[:maxWidth-len(str)]
	ret = append(ret, str)
	return ret
}

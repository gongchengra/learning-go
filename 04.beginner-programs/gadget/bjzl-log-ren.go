package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type MonthDayNum struct {
	month string
	day   string
	num   int
}

func main() {
	oldname := "mysql.log"
	if len(os.Args) > 1 {
		oldname = os.Args[1]
	}
	in, err := os.Open(oldname)
	if err != nil {
		fmt.Println("open file fail:", err)
		os.Exit(-1)
	}
	defer in.Close()
	var mdn []MonthDayNum
	months := []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12"}
	days := []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31"}
	for _, m := range months {
		for _, d := range days {
			mdn = append(mdn, MonthDayNum{month: m, day: d, num: 0})
		}
	}
	br := bufio.NewReader(in)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read err:", err)
			os.Exit(-1)
		}
		data := strings.Split(string(line), "	")
		dates := strings.Split(data[0], "-")
		if len(dates) > 1 {
			month, day := dates[1], dates[2]
			i, _ := strconv.Atoi(data[1])
			for k, v := range mdn {
				if v.month == month && v.day == day {
					mdn[k].num = i
				}
			}
		}
	}
	fmt.Println("日期,1月,2月,3月,4月,5月,6月,7月,8月,9月,10月,11月,12月")
	for _, d := range days {
		fmt.Print(d)
		for _, m := range months {
			for _, v := range mdn {
				if v.month == m && v.day == d {
					fmt.Print(",", v.num)
				}
			}
		}
		fmt.Print("\n")
	}
}

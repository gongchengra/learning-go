package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/golang-module/carbon/v2"
)

func main() {
	start := flag.String("start", "20220513", "Input start date, such as 20220513")
	end := flag.String("end", "20220821", "Input end date, such as 20220822")
	flag.Parse()
	flagset := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { flagset[f.Name] = true })
	startTime := carbon.Parse(*start)
	endTime := carbon.Now()
	if flagset["end"] {
		endTime = carbon.Parse(*end)
	} else {
		if len(os.Args) > 1 {
			endTime = carbon.Parse(os.Args[1])
		}
	}
	fmt.Println(endTime.DiffAbsInDays(startTime))
}

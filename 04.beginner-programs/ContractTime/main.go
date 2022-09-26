package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	//	start:= time.Date(2022, 5, 13, 0, 0, 0, 0, time.UTC)
	start := flag.String("start", "20220513", "Input start date, such as 20220513")
	end := flag.String("end", "20220821", "Input end date, such as 20220822")
	flag.Parse()
	/* The reference time used in time format layouts is the specific time stamp:
	Layout = "01/02 03:04:05PM '06 -0700" // The reference time, in numerical order.
	ANSIC = "Mon Jan _2 15:04:05 2006"
	UnixDate = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822 = "02 Jan 06 15:04 MST"
	RFC822Z = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850 = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123 = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339 = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen = "3:04PM"
	// Handy time stamps.
	Stamp = "Jan _2 15:04:05"
	StampMilli = "Jan _2 15:04:05.000"
	StampMicro = "Jan _2 15:04:05.000000"
	StampNano = "Jan _2 15:04:05.000000000"
	*/
	startTime, err := time.Parse("20060102", *start)
	if err != nil {
		panic(err)
	}
	endTime, err := time.Parse("20060102", time.Now().Format("20060102"))
	if IsFlagPassed("end") {
		endTime, err = time.Parse("20060102", *end)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Start: ", startTime.Format("2006-01-02"))
	fmt.Println("End: ", endTime.Format("2006-01-02"))
	diff := endTime.Sub(startTime)
	fmt.Println("Days: ", diff.Hours()/24)
}

func IsFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

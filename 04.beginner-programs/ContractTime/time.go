package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("The current datetime is:", now)
	fmt.Println("Year:", now.Year())
	fmt.Println("Month:", now.Month())
	fmt.Println("Day:", now.Day())
	fmt.Println("Hour:", now.Hour())
	fmt.Println("Minute:", now.Minute())
	fmt.Println("Second:", now.Second())
	fmt.Println("Nanosecond:", now.Nanosecond())
	fmt.Println("Time: ", now.Format("15:04:05"))
	fmt.Println("Date:", now.Format("Jan 2, 2006"))
	fmt.Println("Timestamp:", now.Format(time.Stamp))
	fmt.Println("ANSIC:", now.Format(time.ANSIC))
	fmt.Println("UnixDate:", now.Format(time.UnixDate))
	fmt.Println("Kitchen:", now.Format(time.Kitchen))
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launch date: %s\n", t.Local())
	utc := now.UTC()
	fmt.Println(utc)
	vals := []string{"2021-07-28", "2020-11-12", "2019-01-05"}
	for _, val := range vals {
		t, err := time.Parse("2006-01-02", val)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(t)
	}
	t1 := now.Add(time.Hour * 27)
	fmt.Println(t1.Format(time.UnixDate))
	t2 := now.AddDate(2, 10, 11)
	fmt.Println(t2.Format(time.UnixDate))
	t3 := now.Add(-time.Hour * 6)
	fmt.Println(t3.Format(time.UnixDate))
	elapsed := t2.Sub(t1)
	fmt.Println(elapsed)
	elapsed = time.Since(t3)
	fmt.Println(elapsed)
	names := []string{
		"Local",
		"UTC",
		"Pacific/Galapagos",
		"Europe/Budapest",
		"Europe/Moscow",
		"Asia/Vladivostok",
		"Antarctica/Vostok",
		"America/New_York",
		"Africa/Tripoli",
	}
	for _, name := range names {
		loc, err := time.LoadLocation(name)
		if err != nil {
			log.Fatal(err)
		}
		t := now.In(loc)
		fmt.Println(loc, ": ", t)
	}
	fmt.Printf("%d\n", time.Now().Unix())
	{
		var t1 = time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
		var t2 = time.Date(2021, time.July, 28, 16, 22, 0, 0, time.UTC)
		var t3 = time.Date(2021, time.July, 28, 16, 22, 0, 0, time.UTC)
		if t1.Equal(t2) {
			fmt.Println("t1 and t2 are equal")
		} else {
			fmt.Println("t1 and t2 are not equal")
		}
		if t2.Equal(t3) {
			fmt.Println("t2 and t3 are equal")
		} else {
			fmt.Println("t2 and t3 are not equal")
		}
		if t1.Before(t2) {
			fmt.Println("t1 is before t2")
		}
		if t3.After(t1) {
			fmt.Println("t3 is after t1")
		}
	}
	{
		v1 := "2022/05/12"
		v2 := "14:55:23"
		v3 := "2014-11-12T11:45:26.37"
		const (
			layout1 = "2006/01/02"
			layout2 = "15:04:05"
			layout3 = "2006-01-02T15:04:05"
		)
		t, err := time.Parse(layout1, v1)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(t.Format(time.UnixDate))
		t, err = time.Parse(layout2, v2)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(t.Format(time.Kitchen))
		t, err = time.Parse(layout3, v3)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(t.Format(time.UnixDate))
	}
	{
		dates := []string{
			"Sat May 28 11:54:40 CEST 2022",
			"Sat May 28 11:54:40 2022",
			"Sat, 28 May 2022 11:54:40 CEST",
			"28 May 22 11:54 CEST",
			"2022-05-28T11:54:40.809289619+02:00",
			"Sat May 28 11:54:40 +0200 2022",
		}
		layouts := []string{
			time.UnixDate,
			time.ANSIC,
			time.RFC1123,
			time.RFC822,
			time.RFC3339Nano,
			time.RubyDate,
		}
		for i := 0; i < len(dates); i++ {
			parsed, err := time.Parse(layouts[i], dates[i])
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(parsed)
		}
	}
	{
		loc, err := time.LoadLocation("Local")
		if err != nil {
			log.Println(err)
		}
		date := "Sat May 28 11:54:40 2022"
		parsed, err := time.ParseInLocation(time.ANSIC, date, loc)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(parsed)
		loc, err = time.LoadLocation("Europe/Moscow")
		if err != nil {
			log.Println(err)
		}
		parsed, err = time.ParseInLocation(time.ANSIC, date, loc)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(parsed)
	}
}

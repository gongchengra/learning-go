package main

import (
	"flag"
	"fmt"
	"testing"
)

var start = flag.String("start", "20220513", "Input start date, such as 20220513")
var end = flag.String("end", "20220821", "Input end date, such as 20220822")

func TestIsFlagPassed(t *testing.T) {
	fmt.Println(*start)
	fmt.Println(*end)
	res := IsFlagPassed("start")
	if res {
		t.Errorf("String mismatch on test")
	}
}

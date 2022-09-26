package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	b := []byte("bool:")
	b = strconv.AppendBool(b, true)
	fmt.Println(string(b))
	b = strconv.AppendQuote(b, `"Fran & Freddie's Diner"`)
	fmt.Println(string(b))
	b = strconv.AppendQuoteRune(b, '☺')
	fmt.Println(string(b))
	b = strconv.AppendQuoteRuneToASCII(b, '☺')
	fmt.Println(string(b))
	b = strconv.AppendQuoteToASCII(b, `"Fran & Freddie's Diner"`)
	fmt.Println(string(b))
	b10 := []byte("int (base 10):")
	b10 = strconv.AppendInt(b10, -42, 10)
	fmt.Println(string(b10))
	b10 = strconv.AppendUint(b10, 42, 10)
	fmt.Println(string(b10))
	b16 := []byte("int (base 16):")
	b16 = strconv.AppendInt(b16, -42, 16)
	fmt.Println(string(b16))
	b16 = strconv.AppendUint(b16, 42, 16)
	fmt.Println(string(b16))
	b32 := []byte("float32:")
	b32 = strconv.AppendFloat(b32, 3.1415926535, 'E', -1, 32)
	fmt.Println(string(b32))
	b64 := []byte("float64:")
	b64 = strconv.AppendFloat(b64, 3.1415926535, 'E', -1, 64)
	fmt.Println(string(b64))
	v := "10"
	if s, err := strconv.Atoi(v); err == nil {
		fmt.Printf("%T, %v", s, s)
	}
	fmt.Println(strconv.CanBackquote("Fran & Freddie's Diner ☺"))
	fmt.Println(strconv.CanBackquote("`can't backquote this`"))
	s := strconv.FormatBool(true)
	fmt.Printf("%T, %v\n", s, s)
	v1 := 3.1415926535
	s32 := strconv.FormatFloat(v1, 'E', -1, 32)
	fmt.Printf("%T, %v\n", s32, s32)
	s64 := strconv.FormatFloat(v1, 'E', -1, 64)
	fmt.Printf("%T, %v\n", s64, s64)
	v2 := int64(-42)
	s10 := strconv.FormatInt(v2, 10)
	fmt.Printf("%T, %v\n", s10, s10)
	s16 := strconv.FormatInt(v2, 16)
	fmt.Printf("%T, %v\n", s16, s16)
	{
		v := uint64(42)
		s10 := strconv.FormatUint(v, 10)
		fmt.Printf("%T, %v\n", s10, s10)
		s16 := strconv.FormatUint(v, 16)
		fmt.Printf("%T, %v\n", s16, s16)
	}
	{
		shamrock := strconv.IsGraphic('☘')
		fmt.Println(shamrock)
		a := strconv.IsGraphic('a')
		fmt.Println(a)
		bel := strconv.IsGraphic('\007')
		fmt.Println(bel)
	}
	{
		c := strconv.IsPrint('\u263a')
		fmt.Println(c)
		bel := strconv.IsPrint('\007')
		fmt.Println(bel)
	}
	{
		i := 10
		s := strconv.Itoa(i)
		fmt.Printf("%T, %v\n", s, s)
	}
	{
		v := "true"
		if s, err := strconv.ParseBool(v); err == nil {
			fmt.Printf("%T, %v\n", s, s)
		}
	}
	{
		v := "3.1415926535"
		if s, err := strconv.ParseFloat(v, 32); err == nil {
			fmt.Printf("%T, %v\n", s, s)
		}
		if s, err := strconv.ParseFloat(v, 64); err == nil {
			fmt.Printf("%T, %v\n", s, s)
		}
		if s, err := strconv.ParseFloat("NaN", 32); err == nil {
			fmt.Printf("%T, %v\n", s, s)
		}
		// ParseFloat is case insensitive
		if s, err := strconv.ParseFloat("nan", 32); err == nil {
			fmt.Printf("%T, %v\n", s, s)
		}
		if s, err := strconv.ParseFloat("inf", 32); err == nil {
			fmt.Printf("%T, %v\n", s, s)
		}
		if s, err := strconv.ParseFloat("+Inf", 32); err == nil {
			fmt.Printf("%T, %v\n", s, s)
		}
		if s, err := strconv.ParseFloat("-Inf", 32); err == nil {
			fmt.Printf("%T, %v\n", s, s)
		}
		if s, err := strconv.ParseFloat("-0", 32); err == nil {
			fmt.Printf("%T, %v\n", s, s)
		}
		if s, err := strconv.ParseFloat("+0", 32); err == nil {
			fmt.Printf("%T, %v\n", s, s)
		}
		//		fmt.Println(string(b), string(b10), string(b16), string(b32), string(b64))
	}
	{
		v32 := "-354634382"
		if s, err := strconv.ParseInt(v32, 10, 32); err == nil {
			fmt.Printf("%T, %v\n", s, s)
		}
		if s, err := strconv.ParseInt(v32, 16, 32); err == nil {
			fmt.Printf("%T, %v\n", s, s)
		}
		v64 := "-3546343826724305832"
		if s, err := strconv.ParseInt(v64, 10, 64); err == nil {
			fmt.Printf("%T, %v\n", s, s)
		}
		if s, err := strconv.ParseInt(v64, 16, 64); err == nil {
			fmt.Printf("%T, %v\n", s, s)
		}
	}
	{
		v := "42"
		if s, err := strconv.ParseUint(v, 10, 32); err == nil {
			fmt.Printf("%T, %v\n", s, s)
		}
		if s, err := strconv.ParseUint(v, 10, 64); err == nil {
			fmt.Printf("%T, %v\n", s, s)
		}
	}
	{
		s := strconv.QuoteRune('☺')
		fmt.Println(s)
	}
	{
		s := strconv.QuoteRuneToASCII('☺')
		fmt.Println(s)
	}
	{
		s := strconv.QuoteRuneToGraphic('☺')
		fmt.Println(s)
		s = strconv.QuoteRuneToGraphic('\u263a')
		fmt.Println(s)
		s = strconv.QuoteRuneToGraphic('\u000a')
		fmt.Println(s)
		s = strconv.QuoteRuneToGraphic('	') // tab character
		fmt.Println(s)
	}
	{
		s := strconv.QuoteToASCII(`"Fran & Freddie's Diner	☺"`)
		fmt.Println(s)
	}
	{
		s := strconv.QuoteToGraphic("☺")
		fmt.Println(s)
		// This string literal contains a tab character.
		s = strconv.QuoteToGraphic("This is a \u263a	\u000a")
		fmt.Println(s)
		s = strconv.QuoteToGraphic(`" This is a ☺ \n "`)
		fmt.Println(s)
	}
	{
		s, err := strconv.Unquote("You can't unquote a string without quotes")
		fmt.Printf("%q, %v\n", s, err)
		s, err = strconv.Unquote("\"The string must be either double-quoted\"")
		fmt.Printf("%q, %v\n", s, err)
		s, err = strconv.Unquote("`or backquoted.`")
		fmt.Printf("%q, %v\n", s, err)
		s, err = strconv.Unquote("'\u263a'") // single character only allowed in single quotes
		fmt.Printf("%q, %v\n", s, err)
		s, err = strconv.Unquote("'\u2639\u2639'")
		fmt.Printf("%q, %v\n", s, err)
	}
	{
		v, mb, t, err := strconv.UnquoteChar(`\"Fran & Freddie's Diner\"`, '"')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("value:", string(v))
		fmt.Println("multibyte:", mb)
		fmt.Println("tail:", t)
	}
	{
		str := "Not a number"
		if _, err := strconv.ParseFloat(str, 64); err != nil {
			e := err.(*strconv.NumError)
			fmt.Println("Func:", e.Func)
			fmt.Println("Num:", e.Num)
			fmt.Println("Err:", e.Err)
			fmt.Println(err)
		}
	}
}

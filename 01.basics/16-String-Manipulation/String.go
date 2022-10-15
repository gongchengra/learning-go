package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
	"unicode"
)

func main() {
	// Creating a string
	str := "Some String"
	sentence := "I'm a long sentence made up of many different words"
	numbers := []string{"one", "two", "three", "four", "five"}
	path := "/home/gabriel/project/"
	// Convert string to lowercase
	lower := strings.ToLower(str)
	fmt.Println(lower)
	// Convert string to uppercase
	upper := strings.ToUpper(str)
	fmt.Println(upper)
	// Check if string contains another string
	if strings.Contains(str, "some") {
		fmt.Println("Yes, the string exists!")
	}
	// Get/Print character range from string
	fmt.Println("Chars 3-10: " + str[3:10])
	fmt.Println("First Five: " + str[:5])
	// Split a string
	words := strings.Split(sentence, " ")
	fmt.Printf("%v \n", words)
	// Split a string on whitespaces using fields
	fields := strings.Fields(sentence)
	fmt.Printf("%v \n", fields)
	// Join an Array of string
	joinStr := strings.Join(numbers, ",")
	fmt.Println(joinStr)
	// Replace (Takes a number of how many replacements it should do -1 = all of them)
	newstr := strings.Replace(str, "Some", "The", -1)
	fmt.Println(newstr)
	// HasPrefix
	isAbsolute := strings.HasPrefix(path, "/")
	fmt.Println(isAbsolute)
	// Has Suffix
	hasTrailingSlash := strings.HasSuffix(path, "/")
	fmt.Println(hasTrailingSlash)
	// Count characters in string
	count := strings.Count(str, "s")
	fmt.Println(count)
	// Dertermine string length
	length := len(str)
	fmt.Println(length)
	fmt.Println(
		// 1
		strings.Index("test", "e"),
		// "a-b"
		strings.Join([]string{"a", "b"}, "-"),
		// == "aaaaa"
		strings.Repeat("a", 5),
		// "bbaa"
		strings.Replace("aaaa", "a", "b", 2),
		// []string{"a","b","c","d","e"}
		strings.Split("a-b-c-d-e", "-"),
	)
	{
		s := "test"
		s = "change me"
		fmt.Println(s)
		for _, c := range s {
			fmt.Printf("%c %[1]d\n", c)
		}
	}
	{
		var buf bytes.Buffer
		buf.WriteString("an ")
		buf.WriteString("old ")
		buf.WriteString("falcon")
		fmt.Println(buf.String())
		w := "falcon"
		fmt.Println(strings.Repeat(w+" ", 5))
	}
	{
		w1 := "falcon"
		w2 := "Falcon"
		if strings.Compare(w1, w2) == 0 {
			fmt.Println("The words are equal")
		} else {
			fmt.Println("The words are not equal")
		}
		if strings.EqualFold(w1, w2) {
			fmt.Println("The words are equal")
		} else {
			fmt.Println("The words are not equal")
		}
		msg := "I saw a fox in the forest. The fox had brown fur."
		output := strings.Replace(msg, "fox", "wolf", 2)
		fmt.Println(output)
		output2 := strings.ReplaceAll(msg, "fox", "wolf")
		fmt.Println(output2)
		fmt.Println(strings.Index(msg, "fox"))
		fmt.Println(strings.LastIndex(msg, "fox"))
	}
	{
		msg := "and old falcon"
		msg2 := "čerešňa"
		fmt.Println(strings.Title(msg))
		fmt.Println(strings.ToUpper(msg))
		fmt.Println(strings.ToUpper(msg2))
		fmt.Println(strings.Title(msg2))
	}
	{
		msg := "a blue 鱼"
		r := '鱼'
		if strings.ContainsRune(msg, r) {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
		fmt.Println("-----------------")
		if strings.Contains(msg, "鱼") {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
	{
		msg := ".an old falcon!"
		cutset := ".!"
		msg2 := strings.Trim(msg, cutset)
		fmt.Println(msg2)
		msg3 := strings.TrimLeft(msg, cutset)
		fmt.Println(msg3)
		msg4 := strings.TrimRight(msg, cutset)
		fmt.Println(msg4)
	}
	{
		msg := "\t\tand old falcon\n"
		fmt.Println(msg)
		msg2 := strings.TrimSpace(msg)
		fmt.Println(msg2)
		msg3 := strings.TrimFunc(msg, func(r rune) bool { return !unicode.IsLetter(r) })
		fmt.Println(msg3)
	}
	{
		msg := "--and old falcon--"
		msg2 := strings.TrimPrefix(msg, "--")
		fmt.Println(msg2)
		msg3 := strings.TrimSuffix(msg, "--")
		fmt.Println(msg3)
	}
	{
		// wget https://raw.githubusercontent.com/janbodnar/data/main/the-king-james-bible.txt -O the-king-james-bible.log
		fileName := "the-king-james-bible.log"
		bs, err := os.ReadFile(fileName)
		if err != nil {
			log.Fatal(err)
		}
		text := string(bs)
		fields := strings.FieldsFunc(text, func(r rune) bool {
			return !('a' <= r && r <= 'z' || 'A' <= r && r <= 'Z' || r == '\'')
		})
		wordsCount := make(map[string]int)
		for _, field := range fields {
			wordsCount[field]++
		}
		keys := make([]string, 0, len(wordsCount))
		for key := range wordsCount {
			keys = append(keys, key)
		}
		sort.Slice(keys, func(i, j int) bool {
			return wordsCount[keys[i]] > wordsCount[keys[j]]
		})
		for idx, key := range keys {
			fmt.Printf("%s %d\n", key, wordsCount[key])
			if idx == 10 {
				break
			}
		}
	}
	{
		builder := strings.Builder{}
		data1 := []byte{72, 101, 108, 108, 111}
		data2 := []byte{32}
		data3 := []byte{116, 104, 101, 114, 101, 33}
		builder.Write(data1)
		builder.Write(data2)
		builder.Write(data3)
		fmt.Println(builder.String())
	}
	{
		t0 := time.Now()
		builder := strings.Builder{}
		for i := 0; i < 100; i++ {
			builder.WriteString("falcon")
		}
		t1 := time.Now()
		result := ""
		for i := 0; i < 100; i++ {
			result += "falcon"
		}
		t2 := time.Now()
		fmt.Println(t1.Sub(t0))
		fmt.Println(t2.Sub(t1))
		fmt.Println(builder.String(), result)
	}
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func isChineseChar(r rune) bool {
	return unicode.Is(unicode.Han, r)
}

func isEnglishChar(r rune) bool {
	return unicode.Is(unicode.Latin, r)
}

func isPunctuationChar(r rune) bool {
	return unicode.Is(unicode.Punct, r) || unicode.Is(unicode.Symbol, r)
}

func isDigitChar(r rune) bool {
	return unicode.Is(unicode.Number, r)
}

func countCharacters(text string) (int, int, int, int) {
	var countChinese, countEnglish, countPunctuations, countDigits int
	for _, r := range text {
		switch {
		case isChineseChar(r):
			countChinese++
		case isEnglishChar(r):
			countEnglish++
		case isPunctuationChar(r):
			countPunctuations++
		case isDigitChar(r):
			countDigits++
		}
	}
	return countChinese, countEnglish, countPunctuations, countDigits
}

func main() {
	var inputReader io.Reader
	var filePath string
	fmt.Println("Please provide the text file path (leave blank to input text):")
	fmt.Scanln(&filePath)
	if filePath != "" {
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()
		inputReader = file
	} else {
		inputReader = os.Stdin
	}
	scanner := bufio.NewScanner(inputReader)
	scanner.Split(bufio.ScanRunes)
	var countChinese, countEnglish, countPunctuations, countDigits int
	for scanner.Scan() {
		text := scanner.Text()
		// Skip EOF character
		if text == string(rune(0x04)) {
			break
		}
		chinese, english, punctuations, digits := countCharacters(text)
		countChinese += chinese
		countEnglish += english
		countPunctuations += punctuations
		countDigits += digits
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	totalCount := countChinese + countEnglish + countPunctuations + countDigits
	fmt.Printf("中文字符：%d\n", countChinese)
	fmt.Printf("英文字符：%d\n", countEnglish)
	fmt.Printf("标点符号：%d\n", countPunctuations)
	fmt.Printf("数字：%d\n", countDigits)
	fmt.Printf("总数：%d\n", totalCount)
}

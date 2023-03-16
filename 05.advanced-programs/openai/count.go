package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// 定义数据结构

type Emoji struct {
	Count int
	Emoji []string
}

func main() {
	// 定义字符串变量
	var text string
	// 读取用户输入的字符串
	reader := bufio.NewReader(os.Stdin)
	for {
		s, _ := reader.ReadString('\n')
		// 如果读取输入字符串为空，则退出循环
		if s == "" {
			break
		}
		text += s
	}
	// 统计字符串字数
	wordLength := len(strings.Fields(text))
	// 统计字符串空格数
	spaceCount := strings.Count(text, " ")
	// 统计字符串标点数
	emoji := regexp.MustCompile(`[,.;'"?!-]`)
	emojiCount := len(emoji.FindAllStringIndex(text, -1))
	// 输出统计结果
	fmt.Printf("字数： %d \n空格数：%d \n标点数：%d\n", wordLength, spaceCount, emojiCount) // 输出结果
}

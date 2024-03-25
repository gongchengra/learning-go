package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Initialize the gestures array
	gestures := []string{"石头", "剪刀", "布"}
	// Initialize the win counter for the player
	wins, lose := 0, 0
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())
	// Display start information message
	fmt.Println("石头剪刀布游戏 Ver.1.00 by H.Yazawa")
	// Indefinite loop for unlimited gameplay
	for {
		// Get the player's gesture through standard input (console)
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("0：石头、1：剪刀、2：布 (输入大于2的数字退出): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		user, err := strconv.Atoi(input)
		if err != nil || user > 2 {
			fmt.Println("Exiting the game.")
			break
		}
		// Decide the computer's gesture using a random number
		computer := rand.Intn(3)
		// Generate the string showing both player's and computer's gestures
		resultString := fmt.Sprintf("玩家：%s、计算机：%s", gestures[user], gestures[computer])
		// Determine the outcome and display the result
		if user == computer {
			fmt.Println(resultString + "...平局！")
		} else if computer == (user+1)%3 {
			fmt.Println(resultString + "...玩家获胜！")
			wins++
		} else {
			fmt.Println(resultString + "...计算机获胜！")
			lose++
		}
	}
	// Display the number of times the player won
	fmt.Printf("玩家获胜次数：%d\n", wins)
	fmt.Printf("计算机获胜次数：%d\n", lose)
}

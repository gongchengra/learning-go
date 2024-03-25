package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	gestures := []string{"石头", "剪刀", "布"}
	wins := [3]int{}
	const playsPerGesture = 10000
	rand.Seed(time.Now().UnixNano())
	fmt.Println("石头剪刀布游戏 - 自动测试")
	for input := 0; input <= 2; input++ {
		for i := 0; i < playsPerGesture; i++ {
			computer := rand.Intn(3)
			if computer == (input+1)%3 {
				wins[input]++
			}
		}
	}
	for input, winCount := range wins {
		fmt.Printf("玩家以 '%s' 总共玩了 %d 次，获胜次数：%d，获胜概率：%.2f%%\n",
			gestures[input], playsPerGesture, winCount, float64(winCount)/float64(playsPerGesture)*100)
	}
}

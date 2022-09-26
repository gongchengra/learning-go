package main

import (
	"math/rand"
	"os"
	"strconv"
	"time"
)

func RandomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func main() {
	pwdLen := 10
	if len(os.Args) > 1 {
		pwdLen, _ = strconv.Atoi(os.Args[1])
	}
	println(RandomString(pwdLen))
	if len(os.Args) > 2 {
		c, _ := strconv.Atoi(os.Args[2])
		for i := 0; i < c; i++ {
			println(RandomString(pwdLen))
		}
	}
}

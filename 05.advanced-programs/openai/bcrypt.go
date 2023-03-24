package main

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"strings"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入密码: ")
	password, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	password = strings.TrimSpace(password) // 去掉换行符和空格
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("加密后的密码:", string(hashedPassword))
}

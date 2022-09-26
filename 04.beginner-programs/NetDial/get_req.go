package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	con, err := net.Dial("tcp", "laogongshuo.com:80")
	checkError(err)
	req := "GET / HTTP/1.0\r\n" +
		"Host: webcode.me\r\n" +
		"User-Agent: Go client\r\n\r\n"
	_, err = con.Write([]byte(req))
	checkError(err)
	res, err := ioutil.ReadAll(con)
	checkError(err)
	fmt.Println(string(res))
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

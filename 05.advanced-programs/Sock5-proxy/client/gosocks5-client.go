package main

import (
	"flag"
	"github.com/go-gost/gosocks5"
	"github.com/go-gost/gosocks5/client"
	"log"
)

var (
	server string
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	flag.StringVar(&server, "p", "", "SOCKS5 server address")
	flag.Parse()
}

func main() {
	addr, err := gosocks5.NewAddr(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	conn, err := client.Dial(server)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	if err := gosocks5.NewRequest(gosocks5.CmdConnect, addr).Write(conn); err != nil {
		log.Fatal(err)
	}
	reply, err := gosocks5.ReadReply(conn)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("reply:", reply)
}

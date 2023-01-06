package main

import (
	"flag"
	"github.com/go-gost/gosocks5/server"
	"log"
	"net"
)

var (
	laddr string
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	flag.StringVar(&laddr, "l", ":1080", "SOCKS5 server address")
	flag.Parse()
}

func main() {
	ln, err := net.Listen("tcp", laddr)
	if err != nil {
		log.Fatal(err)
	}
	srv := &server.Server{
		Listener: ln,
	}
	log.Fatal(srv.Serve(server.DefaultHandler))
}

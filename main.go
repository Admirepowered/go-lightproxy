package main

import (
	"flag"
	"fmt"

	"github.com/Admirepowered/go-lightproxy/core"
)

var addr string
var password string
var server string

func main() {

	addr = "127.0.0.1:6064"
	flag.StringVar(&addr, "l", addr, "bind address in local")
	flag.StringVar(&password, "p", "Google", "password connect to server")
	flag.StringVar(&server, "s", "", "connect server address(user mod should give)")

	flag.Parse()

	if server == "" {
		core.Init("server", addr, password)
		//flag.Usage()

	} else {
		fmt.Println("TTT", server)
		core.Init(server, addr, password)
	}
}

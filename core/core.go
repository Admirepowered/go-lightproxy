package core

import "fmt"

func Init(server string, addr string, password string) {

	fmt.Print("OK")
	if server == "server" {
		Server(addr, password)
	} else {
		Clent(server, addr, password)
	}
}

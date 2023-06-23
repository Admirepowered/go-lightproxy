package core

import "fmt"

func Init(server string, addr string, password string, method string) {

	fmt.Println("OK", method)
	if server == "server" {
		Server(addr, password, method)
	} else {
		Clent(server, addr, password, method)
	}
}

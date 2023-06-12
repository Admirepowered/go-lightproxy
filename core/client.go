package core

import (
	"log"
	"net"

	"github.com/Admirepowered/go-lightproxy/core/loger"
	"github.com/Admirepowered/go-lightproxy/core/parse"
)

func handle_conn(conn net.Conn) {
	headerBuf := make([]byte, 1024)
	_, err := conn.Read(headerBuf)
	if loger.Checkerror(err) {
		return
	}
	if headerBuf[0] == 5 {
		method, remote, port := parse.Paseconnect(conn)
	}
}

func Clent(connet string, addr string, password string) {

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept() //用conn接收链接
		if err != nil {
			log.Fatal(err)
		}
		go handle_conn(conn)
	}

}

package parse

import (
	"net"

	"github.com/Admirepowered/go-lightproxy/core/loger"
)

func Paseconnect(conn net.Conn) (int, int, int) {
	conn.Write([]byte{5, 0})
	headerBuf := make([]byte, 128)
	_, err1 := conn.Read(headerBuf)
	if loger.Checkerror(err1) {
		return 1, -1, -11
	}

	return 0, 0, 0
}

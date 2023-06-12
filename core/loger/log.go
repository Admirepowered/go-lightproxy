package loger

import (
	"fmt"
)

func Checkerror(err error) bool {
	if err != nil {
		fmt.Println("[Error]:", err.Error())
		return true
	}
	return false
}
func Loginfo(a int, info string) {
	if a == 0 {
		fmt.Println("[debug]:", info)
	}

	if a == 1 {
		fmt.Println("[info]:", info)
	}
}

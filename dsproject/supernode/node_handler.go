package main

import (
	"fmt"
	"net"
)

func handleNode(conn net.Conn) {
	fmt.Print(conn.RemoteAddr().String())

	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		fmt.Println(string(buf[0:]))
		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}
	}
}

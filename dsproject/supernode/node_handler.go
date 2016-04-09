package main

import (
	"bufio"
	"dsproject/util"
	"fmt"
	"io"
	"net"
)

func handleNode(conn net.Conn) {
	fmt.Println(conn.RemoteAddr().String())
	reader := bufio.NewReader(conn)
	// Read handler
	for {
		message, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		util.CheckError(err)

		fmt.Println("[Message Received]:" + message)
	}
}

package main

import (
	"bufio"
	"dsproject/util"
	"fmt"
	"io"
	"net"
)

var serverAddr = "127.0.0.1:7070"

func connectServer() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", serverAddr)

	util.CheckError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	util.CheckError(err)

	_, err = conn.Write([]byte("REGISTER SER=666222\r\n"))
	util.CheckError(err)

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

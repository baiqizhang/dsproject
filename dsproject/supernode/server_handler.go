package main

import (
	"bufio"
	"dsproject/util"
	"fmt"
	"io"
	"net"
)

func dialServer() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", util.ServerAddr)
	util.CheckError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	util.CheckError(err)

	_, err = conn.Write([]byte("SUPERNODE REGISTER SN1\r\n"))
	util.CheckError(err)

	reader := bufio.NewReader(conn)
	// Read handler
	for {
		message, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		util.CheckError(err)

		fmt.Println("[Server Message]:" + message)
	}
}

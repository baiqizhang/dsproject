package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

var serverAddr = "127.0.0.1:7070"

func connectServer() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", serverAddr)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	_, err = conn.Write([]byte("test message\r\n"))
	checkError(err)

	//result, err := readFully(conn)
	result, err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Println(string(result))

	os.Exit(0)
}

package main

import (
	"dsproject/util"
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

var serverAddr = "127.0.0.1:7070"

func connectServer() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", serverAddr)

	util.CheckError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	util.CheckError(err)

	_, err = conn.Write([]byte("REGISTER SER=666222\r\n"))
	util.CheckError(err)

	//result, err := readFully(conn)
	result, err := ioutil.ReadAll(conn)
	util.CheckError(err)

	fmt.Println(string(result))

	os.Exit(0)
}

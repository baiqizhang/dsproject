/*
goDaytimeServer
*/
package main

import (
	"dsproject/util"
	"fmt"
	"net"
)

var nodeAddr = make(map[string]string)

func main() {
	// connect to server instance
	go connectServer()

	// listen to node connection requests? (not sure if is required)
	listener, err := net.Listen("tcp", ":6060")
	util.CheckError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	//TODO: add node to set: nodeAddr
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

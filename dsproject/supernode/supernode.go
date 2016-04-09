package main

import (
	"dsproject/util"
	"net"
)

// var nodeAddr = make(map[string]string)

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
		go handleNode(conn)
	}
}

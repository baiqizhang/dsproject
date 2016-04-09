package main

import (
	"dsproject/util"
	"net"
)

func main() {
	// connect to frontend instance
	go dialServer()

	// listen to node connection requests? (not sure if is required)
	listener, err := net.Listen("tcp", ":6060")
	util.CheckError(err)

	for {
		conn, err := listener.Accept()
		util.CheckError(err)

		go handleNode(conn)
	}
}

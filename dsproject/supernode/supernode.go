package main

import (
	"dsproject/util"
	"fmt"
	"net"
)

type client struct {
	conn net.Conn
	name string
}

var clients []client

func main() {
	// connect to frontend instance
	go dialServer()

	// listen to node connection requests? (not sure if is required)
	listener, err := net.Listen("tcp", ":6060")
	util.CheckError(err)
	fmt.Println("Supernode Listening at 6060")
	for {
		conn, err := listener.Accept()
		util.CheckError(err)

		newClient := client{conn: conn, name: "none"}
		clients = append(clients, newClient)
		go handleNode(newClient)
	}
}

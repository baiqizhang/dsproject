package main

import (
	"dsproject/util"
	"fmt"
	"net"
)

var name string
var clients []util.Client

//map for <request id, request struct>
var REQMAP map[string]util.Request = make(map[string]util.Request)

//counter for carnodes which are ordinary nodes and counter for supernodes
var COUNTCAR int = 0
var COUNTSUPER int = 0

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

		newClient := util.Client{Conn: conn, Name: "none"}
		clients = append(clients, newClient)
		go handleNode(newClient)

	}
}

package main

import (
	"dsproject/util"
	"fmt"
	"net"
)

var name string
var clients []util.Client

//REQMAP map for <request id, request struct>
var REQMAP = make(map[string]util.Request)

//COUNTCAR counter for carnodes which are ordinary nodes and counter for supernodes
var COUNTCAR int // 0 is the default value

//COUNTSUPER variable export comment placeholder
var COUNTSUPER int // 0 is the default value

func main() {

	// connect to frontend instance
	go dialServer()

	// listen to peer(SuperNode) connection
	go listenPeer()

	// listen to node connection requests? (not sure if is required)
	listenCarNode()
}

func listenCarNode() {
	listener, err := net.Listen("tcp", ":6060")
	util.CheckError(err)
	fmt.Println("Supernode Listening at 6060 for CarNode connection")
	for {
		conn, err := listener.Accept()
		util.CheckError(err)

		newClient := util.Client{Conn: conn, Name: "none"}
		clients = append(clients, newClient)
		go handleNode(newClient)

	}
}

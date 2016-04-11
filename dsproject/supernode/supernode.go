package main

import (
	"dsproject/util"
	"fmt"
	"net"
)

var name string
var clients []util.Client

var m map[string]util.Client

func main() {
    
    //initialize the map
    m:= make(map[string]util.Client)
    
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

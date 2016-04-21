package main

import (
	"dsproject/util"
	"fmt"
	"os"
)

var name string

//REQMAP map for <request id, request struct>
var REQMAP = make(map[string]util.Request)

//COUNTCAR counter for carnodes which are ordinary nodes and counter for supernodes
var COUNTCAR int // 0 is the default value

//COUNTSUPER variable export comment placeholder
var COUNTSUPER int // 0 is the default value

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Usage: supernode PORT")
		os.Exit(0)
	}

	// connect to frontend instance
	go dialServer(args[0])

	// listen to peer(SuperNode) connection in Ring Topology
	listenPeer(args[0])

	// listen to node connection requests? (not sure if is required)
	// listenCarNode()
}

package main

import (
	"bufio"
	"container/list"
	"dsproject/util"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var curLoc *util.Point
var dest *util.Point
var customerLoc *util.Point

var nodeName string

var myNetAddr string
var idle = false

func main() {
	// Check arguments
	args := os.Args[1:]
	if len(args) != 3 {
		fmt.Println("Usage: carnode XCOORD YCOORD NAME ")
		os.Exit(0)
	}

	// set car current position
	curLoc = util.ParseFloatCoordinates(args[0], args[1])
	if curLoc == nil {
		fmt.Println("incorrect XCOORD format")
		os.Exit(0)
	}

	// Get supernode addresses
	supernodes := getSupernodesAddr()
	for e := supernodes.Front(); e != nil; e = e.Next() {
		fmt.Println("[SuperNode Addr]" + e.Value.(string))
	}
	if supernodes.Len() == 0 {
		fmt.Println("No supernode available, exit")
		os.Exit(0)
	}

	//connect to the first supernode
	dialSuperNode(supernodes.Front().Value.(string))
}

func getSupernodesAddr() list.List {
	var ipList list.List
	conn, err := net.Dial("tcp", util.ServerAddr)
	for err != nil {
		fmt.Println("Unable to connect to the front-end server")
		conn, err = net.Dial("tcp", util.ServerAddr)
		time.Sleep(200 * time.Millisecond)
	}

	conn.Write([]byte("REGISTER CAR1\n"))

	// build a table of supernodes' IP
	connbuf := bufio.NewReader(conn)
	for {
		ip, err := connbuf.ReadString('\n')
		if err != nil || ip == "OK\n" {
			break
		}
		ipList.PushBack(ip[0 : len(ip)-1])
	}
	conn.Close()

	return ipList
}

func dialSuperNode(supernode string) {
	conn, err := net.Dial("tcp", supernode)

	for err != nil {
		conn, err = net.Dial("tcp", supernode)
	}

	connbuf := bufio.NewReader(conn)
	for {
		cmd, _ := connbuf.ReadString('\n')
		if idle == true {
			processCommand(cmd, conn)
		}
	}
}

//processCommand Process commands that's received by car node
func processCommand(cmd string, conn net.Conn) {
	args := strings.Split(cmd, " ")

	if strings.Compare(args[0], "COMPUTE") == 0 {
		point := util.ParseFloatCoordinates(args[1], args[2])
		d := point.DistanceTo(curLoc)

		conn.Write([]byte(myNetAddr + " " + strconv.FormatFloat(d, 'f', 4, 64)))
	} else if strings.Compare(args[0], "PICKUP") == 0 {
		dest = util.ParseFloatCoordinates(args[1], args[2])
	}
}

func driveCustomer(customerLoc *util.Point, dest *util.Point) {
	idle = false

	// simulate picking up customer
	time.Sleep(1500 * time.Millisecond)

	// update current location
	curLoc = customerLoc
	fmt.Println("Customer picked up")

	time.Sleep(1500 * time.Millisecond)
	fmt.Println("Drop customer")
	curLoc = dest
}

package main

import (
	"bufio"
	"container/list"
	"dsproject/util"
	"fmt"
	"log"
	"math"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

//Point p(x,y) struct
type Point struct {
	x float64
	y float64
}

var curLoc Point
var dest Point
var customerLoc Point

var myNetAddr string
var idle bool

func main() {
	fmt.Println("hello world")
	idle = false
	// set car current position
	args := os.Args[1:]
	x, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		fmt.Println("x coordinate is not a valid float value")
		return
	}

	y, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		fmt.Println("y coordinate is not a valid float value")
		return
	}

	curLoc.x = x
	curLoc.y = y

	supernodes := getSupernodesAddr()

	for e := supernodes.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(string))
		//connect to other ip addr
	}

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Cannot start server.")
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn)
		conn.Close()
	}
}

func handleConn(conn net.Conn) {
	connbuf := bufio.NewReader(conn)
	cmd, _ := connbuf.ReadString('\n')

	if idle == true {
		processCommand(cmd, conn)
	}
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

		if err != nil {
			break
		}

		ipList.PushBack(ip[0 : len(ip)-1])
	}
	conn.Close()

	return ipList
}

func waitForRequestThrd(supernode string) {
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
		x, y := parseFloatCoordinates(args[1], args[2])
		d := computeDistance(x, y)

		conn.Write([]byte(myNetAddr + " " + strconv.FormatFloat(d, 'f', 4, 64)))
	} else if strings.Compare(args[0], "PICKUP") == 0 {
		dest.x, dest.y = parseFloatCoordinates(args[1], args[2])
	}
}

func computeDistance(x float64, y float64) float64 {
	return math.Sqrt((curLoc.x-x)*(curLoc.x-x) + (curLoc.y-y)*(curLoc.y-y))
}

func parseFloatCoordinates(strx string, stry string) (float64, float64) {
	x, err := strconv.ParseFloat(strx, 64)
	if err != nil {
		fmt.Println("x coordinate is not a valid float value")
		return math.MaxFloat64, math.MaxFloat64
	}

	y, err := strconv.ParseFloat(stry, 64)
	if err != nil {
		fmt.Println("y coordinate is not a valid float value")
		return math.MaxFloat64, math.MaxFloat64
	}

	return x, y
}

func driveCustomer(customer Point, dest Point) {
	idle = false

	// simulate picking up customer
	time.Sleep(1500 * time.Millisecond)

	// update current location
	curLoc.x = customer.x
	curLoc.y = customer.y
	fmt.Println("Customer picked up")

	time.Sleep(1500 * time.Millisecond)
	fmt.Println("Drop customer")
	curLoc.x = dest.x
	curLoc.y = dest.y
}

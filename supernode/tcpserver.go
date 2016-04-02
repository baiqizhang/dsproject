/*
goDaytimeServer
*/
package main

import (
	"fmt"
	"net"
	"os"
)

//Message base message type
type Message struct {
	src  string
	kind string
	data string
}

var nodeAddr = make(map[string]string)

func main() {
	// service := ":1200"
	// tcpAddr, err := net.ResolveTCPAddr("localhost", service)
	// checkError(err)
	listener, err := net.Listen("tcp", ":8080")
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
		// daytime := time.Now().String()
		// conn.Write([]byte(daytime)) // don't care about return value
		// conn.Close()                // we're finished with this client
	}
}

func handleClient(conn net.Conn) {
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

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

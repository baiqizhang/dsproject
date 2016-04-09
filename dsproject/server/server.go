package main

import (
	"bufio"
	"dsproject/util"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"sync"
)

//
var outgoing = make(chan string)

// Default HTTP Request Handler for UI
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Team 9 %s!</h1>", r.URL.Path[1:])
}

func main() {
	// wait until all routines finish
	var wg sync.WaitGroup
	wg.Add(2)

	//start HTTP UI server at 8080
	go func() {
		defer wg.Done()
		http.HandleFunc("/", defaultHandler)
		http.ListenAndServe(":8080", nil)
	}()

	fmt.Print("web server running on 8080\n")

	// start TCP, listening at 7070
	go func() {
		defer wg.Done()
		listener, err := net.Listen("tcp", ":7070")
		util.CheckError(err)

		for {
			conn, err := listener.Accept()
			if err != nil {
				continue
			}
			go handleClient(conn)
		}
	}()
	// outgoing <- "127.0.0.1:6060\n"
	// println("supernode ip sent")
	wg.Wait()
	// for {
	// 	time.Sleep(1000 * time.Millisecond)
	// 	outgoing <- "test\n"
	// 	println("test message sent")
	// }

}

func handleClient(conn net.Conn) {
	//TODO: add node to set: nodeAddr
	fmt.Println("[New Client]:" + conn.RemoteAddr().String())

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	/*
		// Write handler
		go func() {
			for data := range outgoing {
				writer.WriteString(data)
				writer.Flush()
			}
		}()
	*/

	// Read handler
	for {
		message, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		util.CheckError(err)

		fmt.Print("[Message Received]:" + message)

		words := strings.Split(message, " ")

		// if connection comes from CarNode
		if words[0] == "REGISTER" {
			fmt.Println("[Register]:" + words[1])
			writer.WriteString("127.0.0.1:6060\n")
			writer.WriteString("OK\n")
			writer.Flush()
			conn.Close()
			break
		}
		writer.WriteString("error: message not recognized")
	}
}

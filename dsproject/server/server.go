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

// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Team 9 %s!</h1>", r.URL.Path[1:])
}

func main() {
	// wait until all routines finish
	var wg sync.WaitGroup
	wg.Add(2)

	//start server
	go func() {
		defer wg.Done()
		http.HandleFunc("/", defaultHandler)
		http.ListenAndServe(":8080", nil)
	}()

	fmt.Print("web server running on 8080\n")

	// start TCP
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
	wg.Wait()
}

func handleClient(conn net.Conn) {
	//TODO: add node to set: nodeAddr
	fmt.Print(conn.RemoteAddr().String())

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for {
		message, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		util.CheckError(err)

		fmt.Println("[Message Received]:" + message)

		words := strings.Split(message, " ")

		if words[0] == "REGISTER" {
			fmt.Println("[reg]:" + words[1])
			continue
		}
		writer.WriteString("error: not recognized")
	}
	// var buf [512]byte
	// for {
	// 	_, err := conn.Read(buf[0:])
	// 	if err != nil {
	// 		return
	// 	}
	// 	command := string(buf[0:])
	// 	fmt.Println(command)
	//
	// 	_, err2 := conn.Write(reply)
	// 	if err2 != nil {
	// 		return
	// 	}
	// }
}

package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
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
		checkError(err)

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

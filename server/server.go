package main

import (
	"fmt"
	"net/http"
	"sync"
)

// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Team 9 %s!</h1>", r.URL.Path[1:])
}

func main() {
	// wait until all routines finish
	var wg sync.WaitGroup
	wg.Add(1)

	//start server
	go func() {
		defer wg.Done()
		http.HandleFunc("/", defaultHandler)
		http.ListenAndServe(":8080", nil)
	}()

	fmt.Print("web server running on 8080\n")

	//TODO: start TCP

	wg.Wait()
}

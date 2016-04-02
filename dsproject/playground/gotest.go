package main

import (
	"fmt"
	"net"
	"time"
)

func test() {
	fmt.Println("Welcome to the playground!")
	fmt.Println("The time is", time.Now())

	a := []int{1, 2, 3}
	a = append(a, 5)
	fmt.Println(a)

	ip1 := net.ParseIP("127.0.0.1")
	addr := []net.IP{ip1}

	fmt.Println(addr)
}

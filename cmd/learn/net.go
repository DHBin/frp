package main

import (
	"fmt"
	"net"
)

func LearnNet() {
	url := net.JoinHostPort("127.0.0.1", "8080")
	conn, err := net.Dial("tcp", url)
	if err != nil {
		panic(err)
	}
	_ = conn.Close()
	fmt.Println(url)
}

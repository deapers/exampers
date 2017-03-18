// batchdeal project main.go
package main

import (
	"batchtask"
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	rpc.Register(new(batchtask.BatchTask))

	l, err := net.Listen("tcp", ":9998")
	if err != nil {
		fmt.Printf("Listener tcp err: %s", err)
		return
	}

	for {
		fmt.Println("wating...")
		conn, err := l.Accept()
		if err != nil {
			fmt.Sprintf("accept connection err: %s\n", conn)
		}
		go jsonrpc.ServeConn(conn)
	}
}

package main

import (
	"fmt"
	"net"
	"time"
)

var (
	ln        net.Listener
	closeFlag bool = false
)

func startServer() (err error) {
	ln, err = net.Listen("tcp", ":12345")
	if err != nil {
		return err
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		fmt.Printf("已经开启服务成功哦: %v\n", err)
		if err != nil {
			fmt.Printf("accept error: %v\n", err)
			if closeFlag {
				break
			} else {
				continue
			}
		} else {
			conn.Close()
		}
	}
	return nil
}

func main() {
	go startServer()
	time.Sleep(100 * time.Second)
	closeFlag = true
	ln.Close()
	time.Sleep(1 * time.Second)
}

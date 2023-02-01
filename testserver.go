package main

import (
	"fmt"
	"log"
	"net"

	"github.com/gin-gonic/gin"
)

func handleUDPConnection(conn *net.UDPConn) {
	buffer := make([]byte, 8096)
	n, addr, err := conn.ReadFromUDP(buffer)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("UDP client: ", addr)
		fmt.Println("Received from UDP client: ", string(buffer[:n]))
	}
}

func main() {
	server := gin.Default()
	host, port := "localhost", "41234"
	udpAddr, err := net.ResolveUDPAddr("udp4", fmt.Sprintf("%s:%s", host, port))

	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	for {
		handleUDPConnection(conn)
	}

	fmt.Sprintf("UDP server up and listening on port %s", port)
	server.Run(":8099")
}

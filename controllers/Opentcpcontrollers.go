package controllers

import (
	"bufio"
	"fmt"
	"net"

	"github.com/gin-gonic/gin"
)

func Opentcp(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    200,
		"message": "开启tcp！",
		// "data":    ,
	})
	go Open()
	// listen, err := net.Listen("tcp", "127.0.0.1:20000")
	// if err != nil {
	// 	c.JSON(200, gin.H{
	// 		"code": 200,
	// 		"msg":  "操作失败！",
	// 	})
	// 	return
	// } else {
	// 	fmt.Println("已经开启tcp端口成功！")
	// 	// Okbeng(c)
	// 	for {
	// 		conn, err := listen.Accept() // 建立连接
	// 		if err != nil {
	// 			fmt.Println("accept failed, err:", err)
	// 			continue
	// 		}
	// 		go process(conn) // 启动一个goroutine处理连接
	// 	}

	// }

}

// func Okbeng(c *gin.Context) {
// 	go c.JSON(200, gin.H{
// 		"code":    200,
// 		"message": "服务器端口已开启！",
// 		// "data":    ,
// 	})

// }

func Open() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn) // 启动一个goroutine处理连接
	}
}

func process(conn net.Conn) {
	defer conn.Close() // 关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)
		conn.Write([]byte(recvStr)) // 发送数据
	}
}

package controllers

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"net"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Opentcp(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    200,
		"message": "开启tcp！",
		// "data":    ,
	})
	go Open()

}

// func Closetcp(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"code":    200,
// 		"message": "tcp已经关闭！",
// 		// "data":    ,
// 	})
// 	go Stop()

// }
//
//	func Stop(conn net.Conn) {
//		conn.Close() // 关闭连接
//		// listen, err := net.Listen("tcp", "127.0.0.1:20000")
//		// if err != nil {
//		// 	fmt.Println("listen failed, err:", err)
//		// 	return
//		// }
//		// for {
//		// 	conn, err := listen.Accept() // 建立连接
//		// 	if err != nil {
//		// 		fmt.Println("accept failed, err:", err)
//		// 		continue
//		// 	}
//		// 	go process(conn) // 启动一个goroutine处理连接
//		// }
//	}
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
		// p, _ := strconv.ParseInt(buf[:n] 2, 10)
		// p1, _ := strconv.FormatInt(p, 16)
		recvStr := string(buf[:n])
		if recvStr == "关闭" {
			conn.Close()
		}
		str := hex.EncodeToString(buf[:n]) //转字符串
		dizhi := SubStr(str, 0, 2)
		//bh, _ := hex.DecodeString(dizhi)
		//strconv.ParseInt(str, 16, 64)
		h, _ := strconv.ParseInt(dizhi, 16, 64) //16进制的2位码转10进制的int64位
		//hh, _ := strconv.ParseInt(h, 16, 0)
		//recvstr1 := btox(n)
		// fmt.Println("收到client端发来的数据：", p1)
		// fmt.Println("收到client端发来的数据：", n)
		fmt.Println("收到client端发来的hex数据：", str)
		fmt.Println("收到client端发来的设备地址数据：", dizhi)
		fmt.Println("实际编号数据：", h)
		conn.Write([]byte(recvStr)) // 发送数据
	}
}

// 截取字符串，支持多字节字符
// start：起始下标，负数从从尾部开始，最后一个为-1
// length：截取长度，负数表示截取到末尾
func SubStr(str string, start int, length int) (result string) {
	s := []rune(str)
	total := len(s)
	if total == 0 {
		return
	}
	// 允许从尾部开始计算
	if start < 0 {
		start = total + start
		if start < 0 {
			return
		}
	}
	if start > total {
		return
	}
	// 到末尾
	if length < 0 {
		length = total
	}

	end := start + length
	if end > total {
		result = string(s[start:])
	} else {
		result = string(s[start:end])
	}

	return
}

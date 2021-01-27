package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func processRd(conn net.Conn, message chan<- []byte) {
	defer conn.Close() // 关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		message <- buf[:n]
	}
}

func processWr(conn net.Conn, message <-chan []byte) {
	defer conn.Close()
	wr := bufio.NewWriter(conn)
	for {
		line, ok := <-message
		if !ok {
			log.Printf("write error\n")
			return
		}
		fmt.Print(string(line))
		wr.Write(line)
		wr.Flush()
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		log.Fatal("listen error: %v\n", err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("accept error: %v\n", err)
			continue
		}
		ch := make(chan []byte, 1)
		go processRd(conn, ch)
		go processWr(conn, ch)
	}
}

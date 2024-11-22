package main

import (
	"log"
	"net"
	"time"
)

func do(conn net.Conn){
	defer conn.Close()

	var buf []byte = make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(3*time.Second)

	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello World\r\n"))
}
func main() {
	listener, err := net.Listen("tcp", ":1729")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go do(conn)
	}
}

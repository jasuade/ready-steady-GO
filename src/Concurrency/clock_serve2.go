package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	port := os.Args[1]
	listener, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal(err)
	}

	for {
		con, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(con)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		date, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		log.Print(date)
		if err != nil {
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
}

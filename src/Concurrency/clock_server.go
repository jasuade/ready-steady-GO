package main

import (
	"io"
	"log"
	"net"
	"time"
)

//Server
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
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

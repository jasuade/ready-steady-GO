package main

import (
	"log"
	"net"
)

//Server
func main() {
	listener, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		con, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(con)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	buff := make([]byte, 100)
	c.Read(buff)
	log.Println(string(buff))
}

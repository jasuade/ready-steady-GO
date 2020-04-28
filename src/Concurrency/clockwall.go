package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

//go run clockwall.go Madrid=localhost:8010 Tokio=localhost:8020

func main() {
	input := os.Args[1:]
	var cityAndServer []string
	ch := make(chan string)
	defer close(ch)
	for _, elem := range input {
		cityAndServer = strings.Split(elem, "=")
		go readTime(cityAndServer[0], cityAndServer[1], ch)
		fmt.Println(<-ch)
	}
}

func readTime(city string, server string, ch chan string) {
	conn, err := net.Dial("tcp", server)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	buff := make([]byte, 100)
	conn.Read(buff)

	ch <- fmt.Sprintf("The time in %s is %s\n", city, string(buff))

}

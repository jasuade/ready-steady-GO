package main

import (
	"bufio"
	base64 "encoding/base64"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

const EOS rune = 004
const BUFF_SIZE int = 1024
const ADDR = "localhost"
const PORT = "5000"

func main() {
	server := ADDR
	if len(os.Args) > 1 && os.Args[1] != "" {
		server = os.Args[1]
	}
	port := PORT
	if len(os.Args) > 2 && os.Args[2] != "" {
		port = os.Args[2]
	}

	conn, err := net.Dial("tcp", server+":"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	for {
		msg := receiveMessage(conn)
		msgContent := strings.Split(msg, ";")
		switch msgContent[0] {
		case "file":
			{
				err := createFile(msgContent[1:])
				if err != nil {
					log.Fatal(err)
				}
			}
		case "exit":
			{
				os.Exit(0)
			}
		default:
			{
				fmt.Println(msg)
			}
		}
		sendCommand(conn, os.Stdin)
	}
}

func receiveMessage(conn net.Conn) string {
	buff := make([]byte, BUFF_SIZE)
	var msg string

	for {
		n, err := conn.Read(buff)
		if err != nil {
			log.Fatal(err)
			return ""
		}
		part := string(buff[:n])
		if i := strings.IndexRune(part, EOS); i != -1 {
			msg += part[:i]
			break
		}
		msg += part
	}
	return msg
}

func sendCommand(conn net.Conn, r io.Reader) {
	reader := bufio.NewReader(r)
	text, _ := reader.ReadString('\n')
	conn.Write([]byte(text))
}

func createFile(fileContent []string) error {
	fmt.Println("Content of the file:" + fileContent[0] + fileContent[1])

	f, err := os.Create(fileContent[0])
	if err != nil {
		return err
	}
	defer f.Close()

	decodedContent, _ := base64.StdEncoding.DecodeString(fileContent[1])
	size, err := f.Write(decodedContent)
	if err != nil {
		return err
	}
	fmt.Printf("wrote %d bytes\n", size)
	return nil
}

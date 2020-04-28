package main

import (
	base64 "encoding/base64"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
)

const EOS rune = 004
const BUFF_SIZE int = 1024

func main() {
	listener, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		log.Fatal("Error listening %v ", err)
		os.Exit(1)
	}
	defer listener.Close()
	//defer os.Exit(0)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print("Error executing Acceting request %v ", err)
			continue
		}
		sendMessage(conn, "|| You connected to Jara's FTP server || \n ____________________________________ \n ____________________________________ \n ")
		go handleConn(conn)
	}
}

func sendMessage(conn net.Conn, s string) {
	strings.Trim(s, " ")
	_, err := conn.Write([]byte(s + string(EOS)))
	if err != nil {
		log.Print("Error sending message %v", err)
		return
	}
}

func handleConn(conn net.Conn) {
	buff := make([]byte, 1024)
	for {
		n, err := conn.Read(buff)
		if err != nil && err != io.EOF {
			log.Print("Error %v", err)
			return
		}

		command := fmt.Sprintln(string(buff[:n]))
		strings.Trim(command, " \n")
		s := strings.Split(command, " ")

		switch strings.Trim(s[0], " \n") {
		case "cd":
			changeDirectory(conn, fmt.Sprintf("%s", strings.Trim(s[1], " \n")))
			continue
		case "ls":
			listFiles(conn)
			continue
		case "get":
			getFile(conn, fmt.Sprintf("%s", strings.Trim(s[1], " \n")))
			continue
		case "close":
			closeConnection(conn)
			continue
		default:
			continue
		}
	}

}
func listFiles(conn net.Conn) {
	output, err := exec.Command("ls", "-l").Output()
	log.Println(string(output))
	if err != nil {
		log.Fatal("Error executing ls %v", err)
		return
	}
	sendMessage(conn, fmt.Sprintf("%s", output))
}

func getFile(conn net.Conn, fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	//fileSize := fileInfo.Size()
	fileName = fileInfo.Name()

	buff := make([]byte, BUFF_SIZE)
	var fileContent string

	for {
		n, err := file.Read(buff)
		if err == io.EOF {
			break
		}
		fileContent += string(buff[:n])
	}

	encodedFileContent := base64.StdEncoding.EncodeToString([]byte(fileContent))
	sendMessage(conn, string("file;"+fileName+";"+encodedFileContent))

}

func changeDirectory(conn net.Conn, s string) {
	err := os.Chdir("./" + s)
	if err != nil {
		sendMessage(conn, fmt.Sprintf("Directory %s does not exist", s))
		return
	}
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(dir)
	sendMessage(conn, dir)
}

func closeConnection(conn net.Conn) {
	defer conn.Close()
	sendMessage(conn, "exit")
}

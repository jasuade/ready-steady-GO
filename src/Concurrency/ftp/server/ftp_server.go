package main

import (
	base64 "encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const EOS rune = 004
const BUFF_SIZE int = 1024
const ADDR = "0.0.0.0"
const PORT = "5000"

type Handler struct {
	Addr string
	Port string
	Dir  string
}

func main() {

	args := &Handler{}
	args.readHandlerData()

	listener, err := net.Listen("tcp", args.Addr+":"+args.Port)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer listener.Close()
	//defer os.Exit(0)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error Acceping request", err)
			continue
		}
		err = sendMessage(conn, "|| You connected to Jara's FTP server || \n ____________________________________ \n ____________________________________ \n ")
		if err != nil {
			continue
		}
		go args.handleConn(conn)
	}
}

func (args *Handler) readHandlerData() error {
	args.Dir, _ = filepath.Abs("./")

	if len(os.Args) > 1 && os.Args[1] != "" {
		args.Addr = os.Args[1]
		return nil
	}
	if len(os.Args) > 2 && os.Args[2] != "" {
		args.Port = os.Args[2]
		return nil
	}

	args.Addr = ADDR
	args.Port = PORT
	return nil
}

func sendMessage(conn net.Conn, s string) error {
	strings.Trim(s, " ")
	_, err := conn.Write([]byte(s + string(EOS)))
	if err != nil {
		return err
	}
	return nil
}

func (args *Handler) handleConn(conn net.Conn) {
	buff := make([]byte, BUFF_SIZE)
	for {
		n, err := conn.Read(buff)
		if err != nil && err != io.EOF {
			sendMessage(conn, err.Error())
		}

		commandWithArgs := string(buff[:n])
		command := strings.SplitN(commandWithArgs, " ", 2)

		switch strings.TrimSpace(command[0]) {
		case "cd":
			err := changeDirectory(conn, strings.TrimSpace(command[1]))
			if err != nil {
				sendMessage(conn, err.Error())
			}
			continue
		case "ls":
			listFiles(conn)
			continue
		case "get":
			err := getFile(conn, strings.TrimSpace(command[1]))
			if err != nil {
				sendMessage(conn, err.Error())
			}
			continue
		case "close":
			err := args.closeConnection(conn)
			if err != nil {
				sendMessage(conn, err.Error())
			}
			continue
		case "help":
			sendMessage(conn, "cd <path>\t To move between directories\n"+
				"ls \t\t To list the files and directories\n"+
				"get <file> \t To obtain the indicated file\n"+
				"close\t\t To close connection \n"+
				"help \t\t To display list of commands")
			continue
		default:
			sendMessage(conn, "This is not a valid command, use help to see the valid commands")
			continue
		}
	}
}
func listFiles(conn net.Conn) error {
	output, err := exec.Command("ls", "-l").Output()
	log.Println(string(output))
	if err != nil {
		return err

	}
	sendMessage(conn, fmt.Sprintf("%s", output))
	return nil
}

func getFile(conn net.Conn, fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err

	}
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	fileName = fileInfo.Name()

	buff := make([]byte, BUFF_SIZE)
	var fileContent []byte

	for {
		n, err := file.Read(buff)
		if err == io.EOF {
			break
		}
		fileContent = append(fileContent, buff[:n]...)
	}

	encodedFileContent := base64.StdEncoding.EncodeToString(fileContent)
	sendMessage(conn, string("file;"+fileName+";"+encodedFileContent))
	return nil
}

func changeDirectory(conn net.Conn, s string) error {
	err := os.Chdir("./" + s)
	if err != nil {
		return err
	}
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	log.Println(dir)
	sendMessage(conn, dir)
	return nil
}

func (args *Handler) closeConnection(conn net.Conn) error {
	err := os.Chdir(args.Dir)
	if err != nil {
		return errors.New("Unable to return to original working directory " + err.Error())
	}
	defer conn.Close()
	sendMessage(conn, "exit")
	return nil
}

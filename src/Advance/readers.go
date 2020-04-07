package main

import (
	"fmt"
	"io"
	"strings"
)

//io package specifies Readers interface, to read steams of data
//The Go library includes many implementations of this IF
//This interface has the method Read()
//func (T) Read(b []byte) (n int, err error)
func main() {
	r := strings.NewReader("Hello, Reader!")
	//This example create a strings Reader that consumes 8 bytes each time
	b := make([]byte, 8)
	for {
		n, err := r.Read(b) //This method fulfill the buffer and return the data and an error msg
		fmt.Printf("n = %v err=%v b=%v \n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n]) //It return an error = io.EOF when stream ends
		if err == io.EOF {
			break
		}
	}
}

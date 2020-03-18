package main

import (
	"fmt"
	"time"
)

//You can create your own errors
type MyError struct {
	When time.Time
	What string
}

//Error in another built-in interface
func (e *MyError) Error() string {
	return fmt.Sprintf("At %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(), "'Todo se fue a la puta'",
	}
}

func main() {
	//Functions often return an error value,
	//and calling code should handle errors by testing whether the error equals nil
	if err := run(); err != nil {
		fmt.Println(err) //A nil error denotes success; a non-nil error denotes failure.
	}
}

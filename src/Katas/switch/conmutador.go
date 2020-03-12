package conmutador

import (
	"fmt"
	"runtime"
)

func conmutador() string {
	fmt.Print("GO runs on:  ")
	os := runtime.GOOS
	switch os {
	case "darwin":
		os = "OS x"
		fmt.Print("OS X")
	case "linux":
		os = "Linux"
		fmt.Print("Linux")
	default:
		fmt.Printf("Other OS %s", os)
	}
	return os

}

package main

import (
	"fmt"
	"os"
)

var (
	password = os.Getenv("password")
)
func Cat() string {
	return "miao~~~"
}
func main()  {
	saying := Cat()
	fmt.Print(saying)
}

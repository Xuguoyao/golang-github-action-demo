package main

import "fmt"

var (
	password = "123456"
)
func Cat() string {
	return "miao~~~"
}
func main()  {
	saying := Cat()
	fmt.Print(saying)
}

package main

import "fmt"

func Cat() string {
	return "wang..."
	return "miao~~~"
}
func main()  {
	saying := Cat()
	fmt.Print(saying)
}

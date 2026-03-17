package main

import (
	"fmt"
	"os"
)

func greet(name string) string {
	if name == "" {
		return "Hello, World!"
	}
	return fmt.Sprintf("Hello, %s!", name)
}

func main() {
	name := ""
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	fmt.Println(greet(name))
}

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
	for _, arg := range os.Args[1:] {
		if arg == "--serve" {
			startServer(":8080")
			return
		}
	}
	name := ""
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	fmt.Println(greet(name))
}

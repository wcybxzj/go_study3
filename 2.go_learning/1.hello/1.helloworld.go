package main

import "os"

func main() {
	if len(os.Args) > 1 {
		println("Hello World", os.Args[1])
	}
}




package main

import "os"

func main() {
	print("11111");
	if len(os.Args) > 1 {
		println("Hello World", os.Args[1])
	}
}




package main

import (
	"fmt"

	"github.com/mactsouk/post05"
)

func main() {
	post05.Hostname = "localhost"
	fmt.Println(post05.Port)
	fmt.Println("hello----")
	fmt.Println(post05.Hostname)
}


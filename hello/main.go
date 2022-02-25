package main

import (
	"fmt"

	"github.com/similer/test-repo/greeting"
)

func main() {
	message := greeting.Hello("John")
	fmt.Println(message)
}

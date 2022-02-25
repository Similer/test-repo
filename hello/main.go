package main

import (
	"fmt"

	"github.com/similer/practice-golang/greeting"
)

func main() {
	message := greeting.Hello("John")
	fmt.Println(message)
}

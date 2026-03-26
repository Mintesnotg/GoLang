package main

import (
	"fmt"

	"example/hello/greetings"
)

func main() {
	message := greetings.Greetings("Minte")
	fmt.Println("Hello, World!")
	
	fmt.Println(message)
}

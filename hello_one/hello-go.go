package main

import (
    "fmt"
    "log"
    "example.com/greetings"
)

func main() {
    // Get a greeting message and print it.

    
	log.SetPrefix("greetings: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC)

	message,err := greetings.Greetings("Minte")

	if err != nil { log.Fatal(err) }

	fmt.Println(message)
}
package main

import (
	"fmt"
	"github.com/manthanchauhan/go_getting_started/greetings"
	"log"
)

func main() {
	message, err := greetings.Hello("Manthan")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}

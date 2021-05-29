package main

import (
	"fmt"
	"github.com/manthanchauhan/go_getting_started/greetings"
)

func main() {
	message := greetings.Hello("Manthan")
	fmt.Println(message)
}

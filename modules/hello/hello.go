package main

import (
	"fmt"
	"github.com/valeter/go_tutorial/modules/greetings"
)

func main() {
	message := greetings.Hello("valter")
	fmt.Println(message)
}

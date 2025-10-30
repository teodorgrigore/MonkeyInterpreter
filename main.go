package main

import (
	"MonkeyInterpreter/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language\n", currentUser.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin)
}

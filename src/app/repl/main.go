package main

import (
	"choco/src/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("starting REPL >> %s\n", user.Username)

	fmt.Printf("enter your code\n")

	repl.Start(os.Stdin, os.Stdout)
}

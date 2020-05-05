package main

import (
	"fmt"
	"interpreter/runner"
	"os"
)

func main() {
	filename := ""
	if len(os.Args) == 1 {
		filename = "./main.choco"
	} else {
		filename = os.Args[1]
	}

	_, err := os.Stat(filename)
	if err != nil {
		if len(os.Args) == 1 {
			fmt.Print("[ERROR] filename is not given and couldn't find ./main.choco. please place or pass your .choco file")
		} else {
			fmt.Printf("[ERROR] given filename(%s) is not confirm. please confirm you have correct path", filename)
		}
	}

	fmt.Printf("Running choco...\n")
	fmt.Printf("target file is: %s...\n", filename)

	runner.Run(filename, os.Stdin, os.Stdout)
}

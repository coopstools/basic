package main

import (
	"fmt"
	"github.com/coopstools/basic/repl"
	"os"
	osInfo "os/user"
)

func main() {
	user, err := osInfo.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in  commands\n")
	repl.Start(os.Stdin, os.Stdout)
}

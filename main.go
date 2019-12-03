package main

import (
	"fmt"
	"github.com/MoonShining/monkey-lan/repl"
	"os"
	user2 "os/user"
)

func main() {
	user, err := user2.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is Monkey lan!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}

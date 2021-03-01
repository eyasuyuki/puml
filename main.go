package main

// from book "Writing Interpreter In Go"

import (
	"fmt"
	"github.com/eyasuyuki/puml/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Sprintf("Hello %q! This is the PlantUML lexer!\n", user.Username)
	fmt.Printf("Feel free to type in commands.\n")
	repl.Start(os.Stdin, os.Stdout)
}

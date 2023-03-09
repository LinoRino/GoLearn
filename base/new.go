package base

import (
	"fmt"
	"os"
)

func New(name string, version string) {
	if os.Args[1] == "-v" || os.Args[1] == "version" {
		fmt.Println("v." + version)
		os.Exit(0)
	}
	if os.Args[1] == "-h" || os.Args[1] == "help" {
		fmt.Println("Usage: " + name + " [option] [file | folder]")
	}
}

type Command struct {
	name string
	desc string
	fn func() any
}
package main

import (
	"fmt"
	"solid/components"
	"solid/do"
	"os"
)

func main() {
	// Mesure Time Performance
	if len(os.Args) < 2 {
		fmt.Println(`There is no options`)
		os.Exit(1)
	}
	do.Create()
	do.List()
	if len(os.Args) < 3 {
		fmt.Print(components.Red("Error: "))
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}
}

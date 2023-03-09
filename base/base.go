package base

import (
	"fmt"
	"os"
)

/*Base of cli*/
func Base(name string, explane string, do func()) {
	if os.Args[1] == name {
		if len(os.Args) >= 3 {
			// Show help
			if os.Args[2] == "-h" || os.Args[2] == "--help" {
				fmt.Println("\n" + name + ": " + explane)
				fmt.Println("usage:hello " + name + " <filename string>\n")
				os.Exit(0)
			}
		}
		do()
		os.Exit(0)
	}
}

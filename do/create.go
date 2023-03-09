package do

import (
	"fmt"
	"os"
	"solid/base"
	"solid/components"
	"time"
)

func ErrLog(err error) {
	if err != nil {
		fmt.Println(components.Red("Error: ") + err.Error())
		os.Exit(1)
	}
}

func Create() {
	base.Base("create", "create file", func() {
		var fileName string = ""
		for fileName == "" {
			if len(os.Args) > 2 {
				fileName = os.Args[2]
			} else {
				fileName = components.Txt("Enter the component's name")
			}
		}
		// Measure Time Performance
		now := time.Now()
		file, err := os.Create(fileName+".tsx")
		ErrLog(err)
		_, err = file.WriteString("import { JSX } from \"solid-js\";\n\n")
		ErrLog(err)
		_, err = file.WriteString("export function "+fileName+" (props: any): JSX.Element {\n  return /* JSX here. */\n}")
		ErrLog(err)
		file.Close()
		fmt.Printf("\nFinish create file in %vms.\n", time.Since(now).Milliseconds())
	})
}

package components

import (
	"bufio"
	"fmt"
	"os"
)

func Txt(format string) string {
	result := bufio.NewScanner(os.Stdin)
	fmt.Printf(Blue("? "))
	fmt.Print(format+":  ")
	result.Scan()
	return result.Text()
}
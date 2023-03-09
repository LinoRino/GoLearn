package do

import (
	"fmt"
	"os"
	"solid/base"
	"time"
)

func List() {
	base.Base("list", "check components", func() {
		now := time.Now()
		folder, err := os.Open("src")
		ErrLog(err)
		children, _ := folder.Readdirnames(256)
		for i := 0; i <= len(children)-1; i++ {
			fmt.Println("- " + children[i])
		}
		folder.Close()
		fmt.Printf("\nFinish in %vms", time.Since(now).Milliseconds())

	})
}

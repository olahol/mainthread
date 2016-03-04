package main

import (
	"fmt"
	"github.com/olahol/mainthread"
)

func main() {
	go func() {
		mainthread.Do(func() {
			fmt.Println("Hello from main thread")
		})
		mainthread.Quit()
	}()

	mainthread.Main()
}

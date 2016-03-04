package mainthread

import (
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

// Main thread loop, executes functions.
func Main() {
loop:
	for {
		select {
		case f := <-mainfunc:
			f()
		case <-quit:
			break loop
		}
	}
}

var mainfunc = make(chan func())
var quit = make(chan bool)

// Run function in mainthread.
func Do(f func()) {
	done := make(chan bool, 1)
	mainfunc <- func() {
		f()
		done <- true
	}
	<-done
}

func Quit() {
	quit <- true
}

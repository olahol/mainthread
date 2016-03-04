package mainthread

import (
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

// Main thread loop, executes functions.
func Main() {
	for f := range mainfunc {
		f()
	}
}

var mainfunc = make(chan func())

// Run function in mainthread.
func Do(f func()) {
	done := make(chan bool, 1)
	mainfunc <- func() {
		f()
		done <- true
	}
	<-done
}

# mainthread

[![GoDoc][godoc-img]][godoc]

> Run functions in Go's main thread. Package of Russ Cox's code from [golang-nuts][cox-code]

From [LockOSThread][lockosthread]:

> Some libraries, especially graphical frameworks/libraries like Cocoa,
OpenGL, libSDL all require it's called from the main OS thread or
called from the same OS thread due to its use of thread local data
structures. Go's runtime provides `LockOSThread()` function for this,
but it's notoriously difficult to use correctly.

## Install

```bash
go get github.com/olahol/mainthread
```

## Example

```go
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
```

## [Documentation](godoc)

[godoc]: https://godoc.org/github.com/olahol/go-imageupload
[godoc-img]: https://godoc.org/github.com/olahol/go-imageupload?status.svg
[cox-code]: https://groups.google.com/forum/#!msg/golang-nuts/IiWZ2hUuLDA/SNKYYZBelsYJ
[lockosthread]: https://github.com/golang/go/wiki/LockOSThread

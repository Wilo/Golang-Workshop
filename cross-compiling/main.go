package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello from Go, Operating system: ", runtime.GOOS, " Architecture: ", runtime.GOARCH)
}

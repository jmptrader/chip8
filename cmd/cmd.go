package main

import (
    "github.com/eskrm/chip8"
    "runtime"
)

func init() {
    // This is needed to arrange that main() runs on the main thread
    runtime.LockOSThread()
}

// Command line binary
func main() {
    window := chip8.NewSFMLWindow(64, 32)
    defer window.Release()
    driver := chip8.NewDriver(window)
    driver.Run()
}

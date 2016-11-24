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
    // TODO: Parse command line options
    // Enforce multiples of (64, 32) for window size
    window := chip8.NewSFMLWindow(640, 320)
    defer window.Release()
    driver := chip8.NewDriver(window)
    driver.Run()
}

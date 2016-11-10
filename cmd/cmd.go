package main

import (
    "github.com/eskrm/chip8"
    "runtime"
)

func init() {
    // This is needed to arrange that main() runs on main thread.
    runtime.LockOSThread()
}

// Command line binary
func main() {
    screen := chip8.InitOpenGLScreen()
    defer screen.Release()
    driver := chip8.NewDriver(screen)
    driver.Run()
}

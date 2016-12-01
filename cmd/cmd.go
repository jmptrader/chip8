package main

import (
    "flag"
    "github.com/eskrm/chip8"
    "runtime"
)

func init() {
    // This is needed to arrange that main() runs on the main thread
    runtime.LockOSThread()
}

// Command line binary
func main() {
    width := flag.Uint("width", 640, "the width of the window in pixels")
    height := flag.Uint("height", 320, "the height of the window in pixels")
    romPath := flag.String("rom", "", "the path to a chip8 ROM file")
    flag.Parse()

    window := chip8.NewSFMLWindow(*width, *height)
    defer window.Release()
    driver := chip8.NewDriver(window, *romPath)
    driver.Run()
}

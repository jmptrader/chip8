package chip8

type Window interface {
    Update()
    Draw(x, y uint, sprite []byte)
    Clear()
    ShouldClose() bool
    Release()
}

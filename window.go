package chip8

type Window interface {
    Update()
    Draw(x byte, y byte, sprite []byte)
    ShouldClose() bool
    Release()
}

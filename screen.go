package chip8

type Screen interface {
    Update()
    Draw(x uint, y uint, sprite []byte)
    Clear()
    ShouldClose() bool
    Release()
}

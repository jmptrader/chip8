package chip8

const WindowRight uint = 64
const WindowBottom uint = 32

type Window interface {
    Update()
    Draw(x, y uint, drawable []byte)
    Clear()
    ShouldClose() bool
    Release()
}

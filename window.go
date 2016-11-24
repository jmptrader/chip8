package chip8

const WindowRight uint = 64
const WindowBottom uint = 32

type HexKey byte

type Window interface {
    Update()
    IsKeyPressed(key HexKey) bool
    WaitForKeyPress() HexKey
    Draw(x, y uint, drawable []byte)
    Clear()
    ShouldClose() bool
    Release()
}
